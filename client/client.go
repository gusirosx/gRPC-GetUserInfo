package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "gRPC-GetUserInfo/proto"
)

func main() {
	// Set up a http server.
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintln(ctx.Writer, "Up and running...")
	})

	router.GET("/username/:username", GetUser)

	// Run http server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func GetUser(ctx *gin.Context) {
	// Set up a connection to the server.
	conn, err := Connection()
	if err != nil {
		log.Printf("failed to dial server %s: %v", *serverAddr, err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	userName := ctx.Param("username")
	// Contact the server and print out its response.
	req := &pb.UserRequest{Username: userName}
	response, err := client.GetUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error when calling GetUser": err.Error()})
		log.Fatalln("Error when calling GetUser:", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": response})
}
