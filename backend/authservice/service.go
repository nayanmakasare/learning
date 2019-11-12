package main

import (
	"context"
	"github.com/nayanmakasare/learning/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type authServer struct {
}

func (authServer) Login(ctx context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	return &proto.AuthResponse{}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, authServer{})

	listener, err := net.Listen("tcp", (":5000"))
	if err != nil {
		log.Fatal("Error while Listining:  ", err.Error())
	}
	server.Serve(listener)
}
