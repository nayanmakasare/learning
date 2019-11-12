package main

import "google.golang.org/grpc"

type AuthServer struct {
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, AuthServer{})
}
