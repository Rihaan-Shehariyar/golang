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

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	if strings.Contains(info.FullMethod, "Login") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata missing")
	}

	token := md["Authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "token missing")
	}

	parseToken, err := ValidateToken(token[0])
	if err != nil || !parseToken.Valid {
		return nil, status.Error(codes.Unauthenticated, "metadata missing")
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	newCtx := context.WithValue(ctx, "username", username)

	return handler(newCtx, username)

}
