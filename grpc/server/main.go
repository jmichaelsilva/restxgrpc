package main

import (
	"context"
	"log"
	"net"

	pb "perf_test/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedImageServiceServer
}

func (s *server) SendImage(ctx context.Context, in *pb.SendImageRequest) (*pb.SendImageResponse, error) {
	log.Printf("Received image")
	return &pb.SendImageResponse{Message: "OK"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterImageServiceServer(s, &server{})

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
