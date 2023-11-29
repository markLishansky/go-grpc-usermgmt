package main

import (
	"context"
	"fmt"
	"log"

	// "math/rand"
	"net"
	"os"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	conn *pgx.Conn

	pb.UnimplementedUserManagementServer
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
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
		log.Fatalf("Failed write to file: %v", err)
	}
}

// When user is added, read full userlist from file into
// userlist struct, then append new user and write new userlist back to file
func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	// using postgres
	log.Printf("Recieved: %v", in.GetName())
	createSql := `
	create table if not exists users(
		id SERIAL PRIMARY KEY,
		name text,
		age int
	);
	`

	_, err := s.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	created_user := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
	}

	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}

	_, err = tx.Exec(context.Background(),
		"insert into users(name, age) values($1, $2)", created_user.Name,
		created_user.Age)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return created_user, nil

	// store into users.json
	// var users_list *pb.UserList = &pb.UserList{}
	// var user_id int32 = int32(rand.Intn(100))
	// readBytes, err := os.ReadFile("users.json")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Print("File not found. Creating a new one..")
	// 		UserListMarshal(created_user, users_list)
	// 		return created_user, nil
	// 	} else {
	// 		log.Fatalln("Error reading file: ", err)
	// 	}
	// }

	// if err := protojson.Unmarshal(readBytes, users_list); err != nil {
	// 	log.Fatalf("Failed to parse user list: %v", err)
	// }
	// UserListMarshal(created_user, users_list)
	// return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	// get from the users.json
	// jsonBytes, err := os.ReadFile("users.json")
	// if err != nil {
	// 	log.Fatalf("Failed to read from file: %v", err)
	// }
	// var users_list *pb.UserList = &pb.UserList{}
	// if err := protojson.Unmarshal(jsonBytes, users_list); err != nil {
	// 	log.Fatalf("Unmararshaling failed: %v", err)
	// }

	var users_list *pb.UserList = &pb.UserList{}

	rows, err := s.conn.Query(context.Background(), "select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users_list.Users = append(users_list.Users, &user)
	}

	return users_list, nil
}

func main() {
	database_url := "postgres://postgres:1308@localhost:5432/go_usermgmt"
	conn, err := pgx.Connect(context.Background(), database_url)

	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}

	defer conn.Close(context.Background())
	var user_management_server *UserManagementServer = NewUserManagementServer()
	user_management_server.conn = conn

	if err := user_management_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
