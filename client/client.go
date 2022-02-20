package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	pb "gRPC-GetUserInfo/proto"
)

func main() {

	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
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
