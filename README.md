# gRPC github user API

This project aims to use the GitHub API to get the information from a given account via gRPC protocol and return a JSON Embeded to the client using REST protocol

## How to run this example

1. run grpc server

```sh
$ make run_server
```

2. run gin client

```sh
$ make run_client
```

3. use browser to test it

```sh
$ 'http://localhost:8080/username/gusirosx'
```

## How to create proto files

1. use the makefile

```sh
$ make generate
```