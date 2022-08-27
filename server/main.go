package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "grpclearning/routeguide"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":50051", "listen address")
)

type server struct {
	pb.UnimplementedChecksumServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", *addr)
	s := grpc.NewServer()
	pb.RegisterChecksumServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Server started")
}

func (s *server) ComputeChecksum(ctx context.Context, in *pb.ComputeChecksumRequest) (*pb.ComputeChecksumResponse, error) {
	log.Printf("ComputeChecksum: %v", in.Data)
	return &pb.ComputeChecksumResponse{Checksum: "hello " + in.GetData()}, nil
}
func (s *server) ComputeAddition(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	var salt int32 = 1
	log.Printf("ComputeAddition: %v", in.GetFirst()+in.GetSecond())
	return &pb.ReadResponse{Result: in.GetFirst() + in.GetSecond() + float32(salt)}, nil
}
