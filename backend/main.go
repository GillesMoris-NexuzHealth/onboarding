package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"practice/proto"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedTitleServiceServer
	client *spanner.Client
}

func (s *server) getLastEntry(ctx context.Context) (*proto.LogEntry, error) {
	statement := spanner.Statement{SQL: `SELECT * FROM logEntry ORDER BY created DESC LIMIT 1`}
	iter := s.client.Single().Query(ctx, statement)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var created int64
		var message string
		if err := row.Columns(&created, &message); err != nil {
			return nil, err
		}
		return &proto.LogEntry{Created: created, Message: message}, nil
	}
	return nil, errors.New("No entry found")
}

func (s *server) storeEntry(ctx context.Context, created int64, message string) error {
	fmt.Printf("Storing: '%s' at %d\n", message, created)
	logEntryColumns := []string{"created", "message"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("logEntry", logEntryColumns, []interface{}{created, message}),
	}
	_, err := s.client.Apply(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) Log(ctx context.Context, request *proto.Request) (*proto.LogEntry, error) {
	title := request.GetTitle()
	if title == "last" {
		fmt.Print("Received: 'last'\n")
		lastEntry, err := s.getLastEntry(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Returning last entry at %d: '%s'\n", lastEntry.GetCreated(), lastEntry.GetMessage())
		return lastEntry, nil
	}

	fmt.Printf("Received: '%s'\n", title)
	created := time.Now().UnixMilli()
	message := fmt.Sprintf("Modified: %s", title)
	err := s.storeEntry(ctx, created, message)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Returning: '%s'\n", message)
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
