package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/jaliyaL/loader-api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type loaderServer struct {
	pb.UnimplementedLoaderServiceServer
}

// GenerateLoad generates fake todo tasks
func (s *loaderServer) GenerateLoad(ctx context.Context, req *pb.LoadRequest) (*pb.LoadResponse, error) {
	var todos []*pb.Todo
	for i := 1; i <= int(req.Count); i++ {
		t := &pb.Todo{
			Id:    int32(i),
			Title: fmt.Sprintf("Task %d - %s", i, randomTitle()),
		}
		todos = append(todos, t)
	}
	return &pb.LoadResponse{Todos: todos}, nil
}

func randomTitle() string {
	titles := []string{
		"Write unit tests",
		"Fix bug in auth service",
		"Review PR",
		"Update docs",
		"Refactor API handler",
	}
	rand.Seed(time.Now().UnixNano())
	return titles[rand.Intn(len(titles))]
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoaderServiceServer(grpcServer, &loaderServer{})

	// Enable reflection for grpcurl or other clients
	reflection.Register(grpcServer)

	log.Println("⚡ LoaderService running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
