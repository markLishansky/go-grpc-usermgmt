package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserManagementClient(conn)
	fmt.Println("AAA", c)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)

	new_users["qqq"] = 12
	new_users["marik"] = 2

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("couldn't create user: %v", err)
		}
		log.Printf(`User Details:
NAME: %s
AGE: %v
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("couldn't retrieve users: %v", err)
	}
	log.Print("\nUSER LIST: \n")
	users := r.GetUsers()
	for _, value := range users {
		fmt.Println(value)
	}
}
