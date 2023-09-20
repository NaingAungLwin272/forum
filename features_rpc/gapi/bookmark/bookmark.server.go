package bookmark_gapi

import (
	"github.com/scm-dev1dev5/mtm-community-forum/features_rpc/pb"
	bookmark_service "github.com/scm-dev1dev5/mtm-community-forum/features_rpc/services/bookmark"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookmarkServer struct {
	pb.UnimplementedBookmarkServiceServer
	bookmarkCollection *mongo.Collection
	bookmarkService    bookmark_service.BookmarkService
}

func NewGrpcBookmarkServer(bookmarkCollection *mongo.Collection, bookmarkService bookmark_service.BookmarkService) (*BookmarkServer, error) {
	bookmarkServer := &BookmarkServer{
		bookmarkCollection: bookmarkCollection,
		bookmarkService:    bookmarkService,
	}

	return bookmarkServer, nil
}
