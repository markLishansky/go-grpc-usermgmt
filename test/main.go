package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	marik := &User{
		Name: "marikkk",
		Age:  12,
		Id: 44,
		Followers: &Followers{
			Youtube: 12,
			Vk: 33,
		},
	}

	data, err := proto.Marshal(marik)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(marik.Age)
	fmt.Println(data)

	newMarik := &User{}

	err = proto.Unmarshal(data, newMarik)

	if err != nil {
		log.Fatal(err)
	}


	newMarik.Name = fmt.Sprintf("%s Ebanat", newMarik.Name)
	fmt.Println(newMarik.Followers.GetVk())


}