# gRPC github user API

This project uses the GitHub API to get the information from a given account via gRPC protocol and return a JSON embeded to the client using REST protocol, achieving communication between micro-services in Golang

## How to run this example

1. run grpc server

```sh
$ make run_server
```

2. run gin client

```sh
$ make run_client
```

# Input

1. use browser to test the application using the following link

```sh
http://localhost:8080/username/gusirosx
```

# Output
```
id: 61150315 
name: "Gustavo Rodrigues" 
username: "gusirosx" 
avatarurl: "https://avatars.githubusercontent.com/u/61150315?v=4" 
location: "Uberlândia" 
statistics: {
  followers:6  
  following:6  
  repos:39  
}
listURLs: ["https://api.github.com/users/gusirosx", 
           "https://api.github.com/users/gusirosx/starred", 
           "https://api.github.com/users/gusirosx/repos"]
```

## How to create proto files

1. use the makefile

```sh
$ make generate
```