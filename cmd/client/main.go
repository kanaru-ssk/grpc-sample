package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sethvargo/go-envconfig"
)

type EnvConfig struct {
	Port      int    `env:"PORT,default=8080"`
	ServerUrl string `env:"SERVER_URL,default=server:8080"`
}

func main() {
	ctx := context.Background()

	var envConfig EnvConfig
	if err := envconfig.Process(ctx, &envConfig); err != nil {
		log.Fatalf("failed to get env: %v", err)
	}

	// Set up a connection to the server.
	// conn, err := grpc.NewClient(envConfig.ServerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)

	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		// defer cancel()

		// Use the name from the query parameter if provided, otherwise use the default
		nameToGreet := r.URL.Query().Get("name")

		// s, err := c.SayHello(ctx, &pb.HelloRequest{Name: nameToGreet})
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("could not greet: %v", err), http.StatusInternalServerError)
		// 	return
		// }

		fmt.Fprintf(w, "Greeting: %s", nameToGreet)
	})

	log.Println("Starting HTTP server on", fmt.Sprintf(":%d", envConfig.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", envConfig.Port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
