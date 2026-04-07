package main

import (
	pb "auth/auth/proto"
	"context"
	"errors"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) Login(ctc context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	if req.Username != "admin" || req.Password != "1234" {
		return nil, errors.New("Invalid crednetials")
	}

	token, err := GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
	}, nil

}

func (s *server) GetProfile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	username := ctx.Value("username").(string)

	return &pb.ProfileResponse{
		Username: username,
	}, nil
}
