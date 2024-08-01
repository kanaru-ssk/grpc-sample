package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/kanaru-ssk/grpc-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
	defaultAddr = "server:50051"
)

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// Get the server address from the environment variable or use the default
	addr := os.Getenv("SERVER_URL")
	if addr == "" {
		addr = defaultAddr
	}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
