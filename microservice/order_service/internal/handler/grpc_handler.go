package handler

import (
	"context"
	"fmt"
	"order_service/internal/service"
	pb "shared-proto/order"
	userpb "shared-proto/user"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	service    *service.OrderService
	userClient userpb.UserServiceClient
}

func NewOrderHandler(s *service.OrderService, uc userpb.UserServiceClient) *OrderHandler {
	return &OrderHandler{
		service:    s,
		userClient: uc,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	_, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, fmt.Errorf("User not Found")
	}

	order, err := h.service.CreateOrder(req.UserId, req.Product)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:      order.ID,
			UserId:  order.UserID,
			Product: order.Product,
		},
	}, nil

}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := h.service.GetOrder(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Order: &pb.Order{
			Id:      order.ID,
			UserId:  order.UserID,
			Product: order.Product,
		},
	}, nil
}
