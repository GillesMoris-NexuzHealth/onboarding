package main

import (
	"context"
	"fmt"
	"net"
	"practice/proto"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedTitleServiceServer
	client *spanner.Client
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

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, "projects/prj-nxh-moapr-spanner-dev-8104/instances/csp-moapr-dev/databases/onboarding-gilles", option.WithCredentialsFile("gcloud-credentials.json"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	srv := grpc.NewServer()
	proto.RegisterTitleServiceServer(srv, &server{proto.UnimplementedTitleServiceServer{}, client})
	reflection.Register(srv)

	fmt.Print("Server started, listening.\n")

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
