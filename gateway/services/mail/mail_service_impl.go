package mail_service

import (
	"errors"

	"github.com/gin-gonic/gin"
	mail_client "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients"
	mail_proxy "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail"
	mail_proto "github.com/scm-dev1dev5/mtm-community-forum/gateway/pkg/clients/mail/pb"
)

type MailService struct {
	MailClient mail_client.ServiceClient
}

// ForgetPaswordMail implements MailServiceInterface.
func (mailSvc *MailService) ForgetPasswordMail(ctx *gin.Context) (*mail_proto.MailResponse, error) {
	data := mail_proxy.ForgetPasswordMail(ctx, mailSvc.MailClient.Mail)
	if data == nil {
		return nil, errors.New("failed to send forget password mail")
	}

	return data, nil
}

// SendMail implements NotiServiceInterface.
func (mailSvc *MailService) SendMail(ctx *gin.Context) (*mail_proto.MailResponse, error) {
	data := mail_proxy.SendMail(ctx, mailSvc.MailClient.Mail)
	if data == nil {
		return nil, errors.New("failed to send mail")
	}

	return data, nil
}

func NewMailService(MailClient mail_client.ServiceClient) MailServiceInterface {
	return &MailService{
		MailClient: MailClient,
	}
}
