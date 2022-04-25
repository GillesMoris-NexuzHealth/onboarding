package main

import (
	"context"
	"fmt"
	"net"
	"practice/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedTitleServiceServer
}

func (s *server) Log(ctx context.Context, request *proto.Request) (*proto.LogEntry, error) {
	title := request.GetTitle()

	fmt.Printf("Received: %s\n", title)

	created := time.Now().UnixMilli()
	message := fmt.Sprintf("Modified: %s", title)

	return &proto.LogEntry{Created: created, Message: message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterTitleServiceServer(srv, &server{})
	reflection.Register(srv)

	fmt.Print("Server started, listening.\n")

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
