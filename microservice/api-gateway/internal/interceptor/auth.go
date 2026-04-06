package interceptor

import (
	"api-gateway/internal/middleware"
	"context"
	"strings"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	if strings.Contains(info.FullMethod, "Login") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata missing")
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "token missing")
	}

	tokenStr := authHeader[0]

	token, err := middleware.ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	// inject into context
	ctx = context.WithValue(ctx, "email", email)

	return handler(ctx, req)

}
