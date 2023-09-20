package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/config"
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/gapi"
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/mail_rpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	serverCertFile   = "../cert/server-cert.pem"
	serverKeyFile    = "../cert/server-key.pem"
	clientCACertFile = "../cert/ca-cert.pem"
)

var (
	ctx         context.Context
	mailService services.EmailService
)

func init() {
	mailService = services.NewMailService(ctx)
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// defer mongoclient.Disconnect(ctx)

	startGrpcServer(config)
}

func startGrpcServer(config config.Config) {
	userServer, err := gapi.NewGrpcMailServer(mailService)
	if err != nil {
		log.Fatal("cannot create grpc userServer: ", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		fmt.Println("cannot load TLS credentials: %w", err)
	}

	fmt.Println(tlsCredentials.Info(), "tlsCredentials ....")
	// serverOptions := []grpc.ServerOption{
	// 	grpc.Creds(tlsCredentials),
	// }
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	// grpcServer := grpc.NewServer()

	pb.RegisterMailServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
	fmt.Println("Server listening on port 8080")
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}
