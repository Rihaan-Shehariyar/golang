package main

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// import (
// 	"context"
// 	"strings"

// 	"github.com/golang-jwt/jwt"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/status"
// )

// func AuthInterceptor(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (interface{}, error) {

// 	if strings.Contains(info.FullMethod, "Login") {
// 		return handler(ctx, req)
// 	}

// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Error(codes.Unauthenticated, "metadata missing")
// 	}

// 	values := md.Get("authorization")
// 	if len(values) == 0 {
// 		return nil, status.Error(codes.Unauthenticated, "token missing")
// 	}
// 	tokenStr := values[0]

// 	parseToken, err := ValidateToken(tokenStr)
// 	if err != nil || !parseToken.Valid {
// 		return nil, status.Error(codes.Unauthenticated, "metadata missing")
// 	}

// 	claims := parseToken.Claims.(jwt.MapClaims)
// 	username := claims["username"].(string)

// 	newCtx := context.WithValue(ctx, "username", username)

// 	return handler(newCtx, req)

// }   

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	if strings.Contains(info.FullMethod, "login") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Metadata missing")
	}

	values := md.Get("Authorization")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Metadata missing")
	}

	tokenstr := values[0]

	parsetoken, err := ValidateToken(tokenstr)
	if err != nil || !parsetoken.Valid {
		return nil, status.Error(codes.Unauthenticated, "Metadata missing")
	}

	claims := parsetoken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	newCtx := context.WithValue(ctx, "username", username)

	return handler(newCtx, req)

}
