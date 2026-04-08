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

	if strings.Contains(info.FullMethod, "login") {
		handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missinf")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missinf")
	}

	token := values[0]

	parseToken, err := ValidateToken(token)
	if err != nil || !parseToken.Valid {
		return nil, status.Error(codes.Unauthenticated, "missinf")
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	newCtx := context.WithValue(ctx, "username", username)

	return handler(newCtx, req)

}
