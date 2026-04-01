package handler

import (
	"user_system/internal/service"
	pb "user_system/user/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}
