package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gRPC-GetUserInfo/proto"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50050", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	usr := pb.NewUserServiceClient(conn)

	fmt.Print("Enter the UserName: ")
	var inputUser string
	fmt.Scanln(&inputUser)

	response, err := usr.GetUser(context.Background(), &pb.UserRequest{Username: inputUser})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %v", err)
	}

	log.Printf("Response from server: %v", response)
}
