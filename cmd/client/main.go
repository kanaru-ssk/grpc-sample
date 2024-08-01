package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
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

	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Use the name from the query parameter if provided, otherwise use the default
		nameToGreet := r.URL.Query().Get("name")
		if nameToGreet == "" {
			nameToGreet = *name
		}

		s, err := c.SayHello(ctx, &pb.HelloRequest{Name: nameToGreet})
		if err != nil {
			http.Error(w, fmt.Sprintf("could not greet: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Greeting: %s", s.GetMessage())
	})

	log.Println("Starting HTTP server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
