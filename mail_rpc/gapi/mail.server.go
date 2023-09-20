package gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/services"
)

type mailServer struct {
	pb.UnimplementedMailServiceServer
	mailService services.EmailService
}

func NewGrpcMailServer(mailService services.EmailService) (*mailServer, error) {
	mailServer := &mailServer{
		mailService: mailService,
	}

	return mailServer, nil
}
