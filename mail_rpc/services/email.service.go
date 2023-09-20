package services

import (
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/pb"
)

type EmailService interface {
	ForgetPasswordMail(req *pb.ForgetMailRequest) error
	SendMail(req *pb.MailRequest) error
}
