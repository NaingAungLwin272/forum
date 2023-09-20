package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/auth"
	config "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/config"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/gapi"
	auth_gapi "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/gapi/auth"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/services"
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
	ctx              context.Context
	mongoclient      *mongo.Client
	userService      services.UserService
	deparmentService services.DeparmentService
	teamService      services.TeamService

	userCollection      *mongo.Collection
	deparmentCollection *mongo.Collection
	teamCollection      *mongo.Collection

	jwtManager     *auth.JWTManager
	forgetPassword *auth.JWTManager
	changePassword *auth.JWTManager
	interceptor    *auth.AuthInterceptor
)

func accessibleRoles() map[string][]string {
	const userServicePath = "/pb.UserService/"
	const teamServicePath = "/pb.TeamService/"
	const departmentServicePath = "/pb.DepartmentService/"

	return map[string][]string{
		userServicePath + "CreateUser":       {"manager"},
		userServicePath + "DeleteUser":       {"manager"},
		teamServicePath + "CreateTeam":       {"manager"},
		teamServicePath + "UpdateTeam":       {"manager"},
		teamServicePath + "DeleteTeam":       {"manager"},
		departmentServicePath + "CreateTeam": {"manager"},
		departmentServicePath + "UpdateTeam": {"manager"},
		departmentServicePath + "DeleteTeam": {"manager"},
	}
}

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

	userCollection = mongoclient.Database("mtm_community_forum").Collection("users")
	deparmentCollection = mongoclient.Database("mtm_community_forum").Collection("departments")
	teamCollection = mongoclient.Database("mtm_community_forum").Collection("teams")

	jwtManager = auth.NewJWTManager("thukhaaung")
	forgetPassword = auth.NewJWTManager("thukhaayanchaw")

	userService = services.NewUserService(userCollection, ctx, jwtManager)
	deparmentService = services.NewDeparmentService(deparmentCollection, ctx)
	teamService = services.NewTeamService(teamCollection, ctx)

	interceptor = auth.NewAuthInterceptor(jwtManager, forgetPassword, accessibleRoles())
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
	userServer, err := gapi.NewGrpcUserServer(userCollection, userService)
	deparmentServer, err := gapi.NewGrpcDeparmentServer(deparmentCollection, deparmentService)
	teamServer, err := gapi.NewGrpcTeamServer(teamCollection, teamService)
	authServer, err := auth_gapi.NewGrpcAuthServer(userCollection, jwtManager, forgetPassword, changePassword)

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
	// // }
	// grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		// grpc.UnaryInterceptor(interceptor.Unary()),
		// grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterDeparmentServiceServer(grpcServer, deparmentServer)
	pb.RegisterTeamServiceServer(grpcServer, teamServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

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
