package clients

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	badge_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/badge/pb"
	category_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/category/pb"
	features_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/features/pb"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
	noti_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/noti/pb"
	user_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/user/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServiceClient struct {
	//user
	Auth       user_proto.AuthServiceClient
	User       user_proto.UserServiceClient
	Team       user_proto.TeamServiceClient
	Department user_proto.DeparmentServiceClient
	Category   category_proto.CategoryServiceClient
	Noti       noti_proto.NotiServiceClient
	Mail       mail_proto.MailServiceClient
	Comment    features_proto.CommentServiceClient

	//badges
	Badge     badge_proto.BadgeServiceClient
	UserBadge badge_proto.UserBadgeServiceClient
	UserPoint badge_proto.UserPointServiceClient
	Vote      features_proto.VoteServiceClient
	BookMark  features_proto.BookmarkServiceClient
	Mention   features_proto.MentionServiceClient
	View      features_proto.ViewServiceClient
	Question  features_proto.QuestionServiceClient
}

func InitServiceClient(config *config.Config) (user_proto.AuthServiceClient, user_proto.UserServiceClient, user_proto.TeamServiceClient, user_proto.DeparmentServiceClient, category_proto.CategoryServiceClient, noti_proto.NotiServiceClient, badge_proto.BadgeServiceClient, badge_proto.UserBadgeServiceClient, badge_proto.UserPointServiceClient, mail_proto.MailServiceClient, features_proto.VoteServiceClient, features_proto.BookmarkServiceClient, features_proto.MentionServiceClient, features_proto.ViewServiceClient, features_proto.QuestionServiceClient, features_proto.CommentServiceClient) {
	//users collections
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// grpc.WithTransportCredentials(insecure.NewCredentials())
	transportOption := grpc.WithTransportCredentials(tlsCredentials)

	//badges collections
	badges_connection, berr := grpc.Dial(config.BadgeSvcUrl, transportOption)
	if berr != nil {
		fmt.Println("Could not connect badge:", berr)
	}

	//categories collections
	categories_connection, cerr := grpc.Dial(config.CategorySvcUrl, transportOption)
	if cerr != nil {
		fmt.Println("Could not connect category:", cerr)
	}

	// features collection
	features_connection, ferr := grpc.Dial(config.FeaturesSvcUrl, transportOption)
	if ferr != nil {
		fmt.Println("Could not connect feature:", ferr)
	}

	// mails collection
	mail_connection, merr := grpc.Dial(config.MailSvcUrl, transportOption)
	if merr != nil {
		fmt.Println("Could not connect mail:", merr)
	}

	// noties collection
	noties_connection, nerr := grpc.Dial(config.NotiSvcUrl, transportOption)
	if nerr != nil {
		fmt.Println("Could not connect noti:", nerr)
	}

	//users collections
	users_connection, uerr := grpc.Dial(config.UserSvcUrl, transportOption)
	if uerr != nil {
		fmt.Println("Could not connect user:", uerr)
	}

	// comment_connection, err := grpc.Dial(config.FeaturesSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return user_proto.NewAuthServiceClient(users_connection), user_proto.NewUserServiceClient(users_connection), user_proto.NewTeamServiceClient(users_connection), user_proto.NewDeparmentServiceClient(users_connection), category_proto.NewCategoryServiceClient(categories_connection), noti_proto.NewNotiServiceClient(noties_connection), badge_proto.NewBadgeServiceClient(badges_connection), badge_proto.NewUserBadgeServiceClient(badges_connection), badge_proto.NewUserPointServiceClient(badges_connection), mail_proto.NewMailServiceClient(mail_connection), features_proto.NewVoteServiceClient(features_connection), features_proto.NewBookmarkServiceClient(features_connection), features_proto.NewMentionServiceClient(features_connection), features_proto.NewViewServiceClient(features_connection), features_proto.NewQuestionServiceClient(features_connection), features_proto.NewCommentServiceClient(features_connection)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
