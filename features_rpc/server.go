package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/config"
	bookmark_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/bookmark"
	comment_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/comment"
	mention_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/mention"
	question_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/question"
	view_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/view"
	vote_gapi "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/gapi/vote"
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"

	bookmark_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/bookmark"
	comment_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/comment"
	mention_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/mention"
	question_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/question"
	view_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/view"
	vote_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/vote"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	ctx                context.Context
	mongoclient        *mongo.Client
	questionService    question_service.QuestionService
	questionCollection *mongo.Collection

	commentService    comment_service.CommentService
	commentCollection *mongo.Collection

	voteService    vote_service.VoteService
	voteCollection *mongo.Collection

	viewService    view_service.ViewService
	viewCollection *mongo.Collection

	mentionService     mention_service.MentionService
	mentionCollection  *mongo.Collection
	bookmarkService    bookmark_service.BookmarkService
	bookmarkCollection *mongo.Collection
)

const (
	serverCertFile   = "../cert/server-cert.pem"
	serverKeyFile    = "../cert/server-key.pem"
	clientCACertFile = "../cert/ca-cert.pem"
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoConnection := options.Client().ApplyURI(config.DBUri)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)

	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	questionCollection = mongoClient.Database("mtm_community_forum").Collection("questions")
	questionService = question_service.NewQuestionService(questionCollection, ctx)

	commentCollection = mongoClient.Database("mtm_community_forum").Collection("comments")
	commentService = comment_service.NewCommentService(commentCollection, ctx)

	voteCollection = mongoClient.Database("mtm_community_forum").Collection("votes")
	voteService = vote_service.NewVoteService(voteCollection, ctx)

	viewCollection = mongoClient.Database("mtm_community_forum").Collection("views")
	viewService = view_service.NewViewService(viewCollection, ctx)

	mentionCollection = mongoClient.Database("mtm_community_forum").Collection("mentions")
	mentionService = mention_service.NewMentionService(mentionCollection, ctx)
	bookmarkCollection = mongoClient.Database("mtm_community_forum").Collection("bookmarks")
	bookmarkService = bookmark_service.NewBookmarkService(bookmarkCollection, ctx)
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Couldn't load config ", err)
	}

	defer mongoclient.Disconnect(ctx)

	startGrpcServer(config)
}

func startGrpcServer(config config.Config) {
	questionServer, err := question_gapi.NewGrpcQuestionServer(questionCollection, questionService)
	commentServer, err := comment_gapi.NewGrpcCommentServer(commentCollection, commentService)
	voteServer, err := vote_gapi.NewGrpcVoteServer(voteCollection, voteService)
	viewServer, err := view_gapi.NewGrpcViewServer(viewCollection, viewService)
	mentionServer, err := mention_gapi.NewGrpcMentionServer(mentionCollection, mentionService)
	bookmarkServer, _ := bookmark_gapi.NewGrpcBookmarkServer(bookmarkCollection, bookmarkService)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		fmt.Println("cannot load TLS credentials: %w", err)
	}
	fmt.Println(tlsCredentials.Info(), "tlsCredentials ....")

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	// grpcServer := grpc.NewServer()

	pb.RegisterQuestionServiceServer(grpcServer, questionServer)
	pb.RegisterCommentServiceServer(grpcServer, commentServer)
	pb.RegisterVoteServiceServer(grpcServer, voteServer)
	pb.RegisterViewServiceServer(grpcServer, viewServer)
	pb.RegisterMentionServiceServer(grpcServer, mentionServer)
	pb.RegisterBookmarkServiceServer(grpcServer, bookmarkServer)

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
