package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	config "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/config"
	bageGapi "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/gapi/badge_api"
	userBadgeGapi "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/gapi/user_badge_api"
	UserPointGapi "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/gapi/user_point_api"

	"github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/pb"
	badgeServices "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/badge_service"
	userBadgeServices "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/user_badges_service"
	UserPointServices "github.com/scm-dev1dev5/mtm-community-forum/badges_rpc/services/user_point_service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	ctx                 context.Context
	mongoclient         *mongo.Client
	badgeService        badgeServices.BadgeService
	badgeCollection     *mongo.Collection
	userPointService    UserPointServices.UserPointService
	userPointCollection *mongo.Collection
	userBadgeCollection *mongo.Collection
	userBadgeService    userBadgeServices.UserBadgeService
	commentCollection   *mongo.Collection
	userCollection      *mongo.Collection
)

func init() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB successfully connected...")

	badgeCollection = mongoclient.Database("mtm_community_forum").Collection("badges")
	badgeService = badgeServices.NewBadgeService(badgeCollection, ctx)

	userBadgeCollection = mongoclient.Database("mtm_community_forum").Collection("user_badges")
	userBadgeService = userBadgeServices.NewUserBadgeService(userBadgeCollection, ctx)

	commentCollection = mongoclient.Database("mtm_community_forum").Collection("comments")

	userPointCollection = mongoclient.Database("mtm_community_forum").Collection("user_points")

	userCollection = mongoclient.Database("mtm_community_forum").Collection("users")
	userPointService = UserPointServices.NewUserPointService(
		userPointCollection,
		userBadgeCollection,
		badgeCollection,
		commentCollection,
		userCollection, ctx)

}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	startGrpcServer(config)
}

func startGrpcServer(config config.Config) {
	UserPointServer, err := UserPointGapi.NewGrpcUserPointServer(userPointCollection, userPointService)
	if err != nil {
		log.Fatal("cannot create grpc userServer: ", err)
	}

	badgeServer, err := bageGapi.NewGrpcBadgeServer(badgeCollection, badgeService)
	if err != nil {
		log.Fatal("cannot create grpc userServer: ", err)
	}

	userBadgeServer, err := userBadgeGapi.NewGrpcUserBadgeServer(userBadgeCollection, userBadgeService)
	if err != nil {
		log.Fatal("cannot create grpc userServer: ", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		fmt.Println("cannot load TLS credentials: %w", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	// grpcServer := grpc.NewServer()

	pb.RegisterBadgeServiceServer(grpcServer, badgeServer)
	pb.RegisterUserPointServiceServer(grpcServer, UserPointServer)
	pb.RegisterUserBadgeServiceServer(grpcServer, userBadgeServer)

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
