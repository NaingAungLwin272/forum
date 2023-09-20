package middleware

import (
	"context"
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/scm-dev1dev5/mtm-community-forum/noti_rpc/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func validateToken(tokenString string) error {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	secret := config.JWT_SECRET
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}
	return nil
}

func extractTokenFromMetadata(md metadata.MD) (string, error) {
	tokenString := ""
	if tokenValues := md.Get("authorization"); len(tokenValues) > 0 {
		tokenString = strings.TrimPrefix(tokenValues[0], "Bearer ")
	}

	if tokenString == "" {
		return "", status.Errorf(codes.Unauthenticated, "authorization token not found")
	}

	return tokenString, nil
}

func extractAndValidateToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata not found")
	}

	tokenString, err := extractTokenFromMetadata(md)
	if err != nil {
		return "", err
	}

	if err := validateToken(tokenString); err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	_, err := extractAndValidateToken(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func JwtMiddlewareStream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	_, err := extractAndValidateToken(ss.Context())
	if err != nil {
		return err
	}

	err = handler(srv, ss)
	if err != nil {
		log.Printf("gRPC error: %s", err.Error())
	}

	return err
}
