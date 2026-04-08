package main

import (
	pb "auth/user/proto"
	"context"
	"errors"
)

type server struct {
	pb.UnimplementedLoginServiceServer
}

func (s *server) LoginRequest(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	if req.Username != "" || req.Password != "1234" {
		return nil, errors.New("Invalid Credentials")
	}

	token, err := GenerateToken(req.Username)
	if err != nil {
		return nil, errors.New("Eror")
	}

	return &pb.LoginResponse{Token: token}, nil

}


