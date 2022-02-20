package main

import (
	"encoding/json"
	"fmt"
	pb "gRPC-user/proto"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement Service Server
type server struct {
	pb.UnimplementedUserServiceServer
}

// Struct used to embed the response
type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"login"`
	AvatarURL  string `json:"avatar_url"`
	Location   string `json:"location"`
	Followers  int64  `json:"followers"`
	Following  int64  `json:"following"`
	Repos      int64  `json:"public_repos"`
	Gists      int64  `json:"public_gists"`
	URL        string `json:"url"`
	StarredURL string `json:"starred_url"`
	ReposURL   string `json:"repos_url"`
}

// GetUser implements server to get the user on Github
func (s *server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	usr := User{}
	jsonErr := json.Unmarshal(body, &usr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &pb.UserResponse{
		Id:        usr.ID,
		Name:      usr.Name,
		Username:  usr.Username,
		Avatarurl: usr.AvatarURL,
		Location:  usr.Location,
		Statistics: &pb.Statistics{
			Followers: usr.Followers,
			Following: usr.Following,
			Repos:     usr.Repos,
			Gists:     usr.Gists,
		},
		ListURLs: []string{usr.URL, usr.StarredURL, usr.ReposURL},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen on port 50050: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
