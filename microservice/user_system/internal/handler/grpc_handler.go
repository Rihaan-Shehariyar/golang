package handler

import (
	"context"
	"fmt"
	pb "shared-proto/user"
	"user_system/internal/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	user, err := h.service.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil

}

func (h *UserHandler) Login(ctc context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	_, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid Credentials")
	}

	return &pb.LoginResponse{
		Token: "valid-Token",
	}, nil

}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {

	user, err := h.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		User: &pb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil

}

func (h *UserHandler) ListUser(ctc context.Context, _ *pb.Empty) (*pb.ListUsersResponse, error) {

	users, err := h.service.ListUser()

	if err != nil {
		return nil, err
	}

	var pbUser []*pb.User
	for _, u := range users {
		pbUser = append(pbUser, &pb.User{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return &pb.ListUsersResponse{Users: pbUser}, nil

}
