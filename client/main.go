package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"

	pb "grpclearning/routeguide"
)

var (
	addr   = flag.String("addr", ":50051", "listen address")
	data   = flag.String("data", "", "Data file to read from")
	first  = flag.Int64("first", 0, "First number to add")
	second = flag.Int64("second", 0, "Second number to add")
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(local.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChecksumClient(conn)
	r, err := c.ComputeChecksum(ctx, &pb.ComputeChecksumRequest{Data: *data})
	if err != nil {
		log.Fatalf("could not compute checksum: %v", err)
	}
	log.Printf("Checksum: %x", r.Checksum)

	read, err := c.ComputeAddition(ctx, &pb.ReadRequest{First: float32(*first), Second: float32(*second)})
	if err != nil {
		log.Fatalf("could not compute addition: %v", err)
	}
	log.Printf("Addition: %v", read.Result)
}
