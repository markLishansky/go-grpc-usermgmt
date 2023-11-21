package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())

	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(),
		Age: in.GetAge(),
		Id: user_id,
	}
	s.user_list.Users = append(s.user_list.Users, created_user)
	return created_user, nil
}

func (s *UserManagementServer) GetUsers (ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}


func main() {
	var user_management_server *UserManagementServer = NewUserManagementServer()
	if err := user_management_server.Run(); err != nil {
		log.Fatalf("failed to cerver: %v", err)
	}
}