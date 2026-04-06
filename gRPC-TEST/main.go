package main

import (
	"context"
	"log"
	"net"
	pb "test/proto/user"

	"google.golang.org/grpc"
)

// type server struct {
// 	pb.UnimplementedHelloServiceServer
// }

// func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

// 	log.Println("Recieved:", req.Name)

// 	return &pb.HelloResponse{
// 		Name: req.Name,
// 	}, nil

// }

// func main() {

// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	grpcServer := grpc.NewServer()

//  pb.RegisterHelloServiceServer(grpcServer,&server{})
//  log.Println("Server Running on port : 50051")
//  if err := grpcServer.Serve(lis);err!=nil{
//   log.Fatal(err)
// }

// }

// type server struct {
// 	pb.UnimplementedHelloServiceServer
// 	users  map[int32]*pb.User
// 	nextID int32
// }

// func newServer() *server {
// 	return &server{
// 		users:  make(map[int32]*pb.User),
// 		nextID: 1,
// 	}
// }

// func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

// 	user := &pb.User{
// 		Id:    s.nextID,
// 		Name:  req.Name,
// 		Email: req.Email,
// 		Age:   req.Age,
// 	}

// 	s.users[s.nextID] = user
// 	s.nextID++

// 	return &pb.UserResponse{User: user}, nil

// }

// func (s *server) GetUser(ctc context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {

// 	user, ok := s.users[req.Id]
// 	if !ok {
// 		return nil, nil
// 	}

// 	return &pb.UserResponse{User: user}, nil

// }

// func (s *server) ListUsers(ctc context.Context, _ *pb.Empty) (*pb.UserListResponse, error) {
// 	var user []*pb.User

// 	for _, u := range s.users {
// 		user = append(user, u)
// 	}

// 	return &pb.UserListResponse{User: user}, nil

// }

// func main() {

// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	grpcServer := grpc.NewServer()

// 	s := newServer()

// 	pb.RegisterHelloServiceServer(grpcServer, s)

// 	log.Println("Server Running on 50051:")

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatal(err)
// 	}

// }

type server struct {
	pb.UnimplementedHelloServiceServer
	users  map[int32]*pb.User
	nextId int32
}

func newServer() *server {
	return &server{
		users:  make(map[int32]*pb.User),
		nextId: 1,
	}
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	user := &pb.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	s.users[s.nextId] = user
	s.nextId++

	return &pb.UserResponse{User: user}, nil

}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, ok := s.users[req.Id]
	if !ok {
		return nil, nil
	}

	return &pb.UserResponse{User: user}, nil

}

func (s *server) ListUsers(ctx context.Context, _ *pb.Empty) (*pb.UserListResponse, error) {

	var users []*pb.User

	for _, u := range s.users {
		users = append(users, u)
	}

	return &pb.UserListResponse{User: users}, nil

}

func main() {

	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	g := newServer()

	pb.RegisterHelloServiceServer(grpcServer, g)

	grpcServer.Serve(lis)

}
