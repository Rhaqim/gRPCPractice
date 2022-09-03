package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "grpclearning/checks"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
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
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	s := grpc.NewServer(grpc.Creds(altsTC))
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
