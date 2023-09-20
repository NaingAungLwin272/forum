package gapi

import (
	"context"
	"log"

	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (mailServer *mailServer) ForgetPasswordMail(ctx context.Context, req *pb.ForgetMailRequest) (*pb.MailResponse, error) {
	log.Println("mail is here")
	if mailServer.mailService == nil {
		return nil, status.Error(codes.Internal, "mailService is not initialized")
	}
	err := mailServer.mailService.ForgetPasswordMail(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.MailResponse{
		IsSuccess: true,
	}

	return res, nil
}

func (mailServer *mailServer) SendMail(ctx context.Context, req *pb.MailRequest) (*pb.MailResponse, error) {
	if mailServer.mailService == nil {
		return nil, status.Error(codes.Internal, "mailService is not initialized")
	}
	err := mailServer.mailService.SendMail(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.MailResponse{
		IsSuccess: true,
	}

	return res, nil
}
