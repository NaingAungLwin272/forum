package comment_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	comment_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/comment"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentServer struct {
	pb.UnimplementedCommentServiceServer
	commentCollection *mongo.Collection
	commentService    comment_service.CommentService
}

func NewGrpcCommentServer(commentCollection *mongo.Collection, commentService comment_service.CommentService) (*CommentServer, error) {
	commentServer := &CommentServer{
		commentCollection: commentCollection,
		commentService:    commentService,
	}

	return commentServer, nil
}
