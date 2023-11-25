package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)



func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}


type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
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

func UserListMarshal(created_user *pb.User, users_list *pb.UserList) {
	users_list.Users = append(users_list.Users, created_user)
	jsonBytes, err := protojson.Marshal(users_list)
	if err != nil {
		log.Fatalf("JSON Marshalling failed: %v", err)
	}
	if err := os.WriteFile("users.json", jsonBytes, 0664); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())
	readBytes, err := os.ReadFile("users.json")
	var users_list *pb.UserList = &pb.UserList{}
	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(),
		Age: in.GetAge(),
		Id: user_id,
	}
	if err != nil {
		if os.IsNotExist(err) {
			log.Print("File not found. Creating a new one..")
			UserListMarshal(created_user, users_list)
			return created_user, nil
		} else {
			log.Fatalln("Error reading file: ", err)
		}
	}

	if err := protojson.Unmarshal(readBytes, users_list); err != nil {
		log.Fatalf("Failed to parse user list: %v", err)
	}
	UserListMarshal(created_user, users_list)
	return created_user, nil
}

func (s *UserManagementServer) GetUsers (ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	jsonBytes, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalf("Failed to read from file: %v", err)
	}
	var users_list *pb.UserList = &pb.UserList{}
	if err := protojson.Unmarshal(jsonBytes, users_list); err != nil {
		log.Fatalf("Unmararshalong failed: %v", err)
	}

	return users_list, nil
}



func main() {
	var user_management_server *UserManagementServer = NewUserManagementServer()
	if err := user_management_server.Run(); err != nil {
		log.Fatalf("failed to cerver: %v", err)
	}
}