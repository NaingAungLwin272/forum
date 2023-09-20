package services

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strings"

	config "github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/config"
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/pb"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailServiceImpl struct {
	ctx context.Context
}

func NewMailService(ctx context.Context) EmailService {
	return &MailServiceImpl{ctx}
}

// ForgetPaswordMail implements EmailService.
func (*MailServiceImpl) ForgetPasswordMail(req *pb.ForgetMailRequest) error {
	email := strings.ToLower(req.Email)
	email_regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !email_regex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	from := mail.NewEmail(config.SEND_FROM_NAME, config.SEND_FROM_ADDRESS)
	subject := "Reset your password for mtm-community-forum"
	to := mail.NewEmail(email, email)
	// plainTextContent := "Your reset password link"

	// Read the template file
	// forgetPassword := "services/forget_password.html"
	forgetPassword := "services/forget_password_2.html"
	// address := to.Address;
	// htmlContent := fmt.Sprintln("Hello, <br>

	// Follow this link to reset your mtm-community-forum password for your", address ,"account. <br><br>

	// <strong>https://example.com</strong> <br><br>

	// If you didn’t ask to reset your password, you can ignore this email. <br>

	// Thanks, <br><br>

	// <strong>Your mtm-community-forum team<strong>");

	tpl, err := template.ParseFiles(forgetPassword)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	// Define the data for the template
	// data := struct {
	// 	Email  string
	// 	Token  string
	// 	Origin string
	// 	Name   string
	// }{
	// 	Email:  email,
	// 	Token:  req.Token,
	// 	Origin: *req.Origin,
	// 	Name:   req.Name,
	// }
	data := struct {
		Email  string
		Token  string
		Origin string
		Name   string
		FRONT_END_URL string
		FRONT_ADMIN_URL string
	}{
		Email:  email,
		Token:  req.Token,
		Origin: *req.Origin,
		Name:   req.Name,
		FRONT_END_URL: config.FRONT_END_URL,
		FRONT_ADMIN_URL: config.FRONT_ADMIN_URL,
	}

	// Create a buffer to store the rendered HTML template
	var emailBody bytes.Buffer
	err = tpl.Execute(&emailBody, data)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}

	content := mail.NewContent("text/html", emailBody.String())
	message := mail.NewV3MailInit(from, subject, to, content)
	client := sendgrid.NewSendClient(config.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		log.Fatal(err)
	}

	// Check if it was sent
	statusCode := response.StatusCode
	fmt.Println(statusCode, "statusCode........")
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println("Email sent!")
	}
	return err
}

// SendMail implements EmailService.
func (*MailServiceImpl) SendMail(req *pb.MailRequest) error {
	email := strings.ToLower(req.Email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	if req.Subject == "" {
		return fmt.Errorf("subject can't be empty")
	}
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	from := mail.NewEmail(config.SEND_FROM_NAME, config.SEND_FROM_ADDRESS)
	to := mail.NewEmail(email, email)
	link := req.Link
	logInLink := fmt.Sprintf("%s/signin", config.FRONT_END_URL)
	subj := strings.Split(req.Subject, "-")[0]
	displayName := strings.Split(req.Subject, "-")[1]
	plainTextContent := ""
	htmlContent := ""

	switch req.Type {
	case 1:
		// Level up
		plainTextContent = "You have reached a new level"
		htmlContent = fmt.Sprintf(`Congratulations! %s , <br>
		 You've Reached a New Level in mtm-community-forum! <br><br>
		 Follow the link to view your level <br> %s <br><br>
		 We look forward to seeing your continued progress and achievements as you continue to explore and share your expertise. <br>
		 Thank you for being an integral part of the mtm-community-forum. <br><br>
		 Best regards,<br>
		 Your MTM Community Forum Team`, displayName, *link)
	case 2:
		// Reply
		plainTextContent = "You have received a reply on your post"
		htmlContent = fmt.Sprintf(`Hello %s , <br>
		You've received a reply on your post in mtm-community-forum! <br><br>
		Follow the link to view your reply <br> %s <br><br>
		Make sure to check it out and engage with the discussion. Your active participation helps create a vibrant community. <br>
		Thank you for being a valued member of the mtm-community-forum. <br><br>
		Best regards,<br>
		Your MTM Community Forum Team`, displayName, *link)
	case 3:
		// Mention
		plainTextContent = "You have been mentioned in a post"
		htmlContent = fmt.Sprintf(`Hello %s , <br>
		You have been mentioned in a post on mtm-community-forum! <br><br>
		Follow the link to view mentioned post <br>
		%s <br><br>
		Make sure to check it out and engage with the discussion. Your active participation helps create a vibrant community. <br>
		Thank you for being a valued member of the mtm-community-forum. <br><br>
		Best regards,<br>
		Your MTM Community Forum Team`, displayName, *link)
	case 4:
		// Solved
		plainTextContent = "Your comment has been marked as the solved solution"
		htmlContent = fmt.Sprintf(`Hello %s , <br>
		Your comment has been marked as the solved solution on the post in mtm-community-forum! <br><br>
		Follow the link to view solved post <br>
		%s <br><br>
		Thank you for your contribution and providing a valuable solution to the community. <br>
		Keep up the great work! <br><br>
		Best regards,<br>
		Your MTM Community Forum Team`, displayName, *link)
	case 5:
		// Voted
		plainTextContent = "Your post has received a new vote"
		htmlContent = fmt.Sprintf(`Hello %s , <br>
		Your post in mtm-community-forum has received a new vote! <br><br>
		Follow the link to view voted post <br>
		%s <br><br>
		Thank you for sharing valuable content with the community. Your contributions are appreciated! <br><br>
		Best regards,<br>
		Your MTM Community Forum Team`, displayName, *link)
	case 6:
		// System message
		plainTextContent = "System Message"
		htmlContent = fmt.Sprintf(`Hello! %s , <br>
		You have received a system message from mtm-community-forum! <br><br>
		Follow the link to view message: <br><br>
		%s <br><br>
		Please review the message and take any necessary actions accordingly. <br><br>
		Best regards,<br>
		Your MTM Community Forum Team`, displayName, *link)
	case 7:
		plainTextContent = "Login Information"
		displayName := displayName
		emailAddress := to.Address
		logInLink := logInLink
		tmpl, err := template.ParseFiles("services/welcome_email.html")
		if err != nil {
			log.Fatal("Error parsing template:", err)
		}
		data := struct {
			DisplayName  string
			EmailAddress string
			LogInLink    string
		}{
			DisplayName:  displayName,
			EmailAddress: emailAddress,
			LogInLink:    logInLink,
		}

		// Execute the template with the data
		var emailBody bytes.Buffer
		err = tmpl.Execute(&emailBody, data)
		if err != nil {
			log.Fatal("Error executing template:", err)
		}
		htmlContent = emailBody.String()
	default:
		// Invalid type
		return fmt.Errorf("invalid mail type. Please provide  type from 1 to 7")
	}

	message := mail.NewSingleEmail(from, subj, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		log.Fatal(err)
	}

	// Check if it was sent
	statusCode := response.StatusCode
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println("Email sent!")
	}
	return nil
}
