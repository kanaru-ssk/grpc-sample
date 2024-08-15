package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sethvargo/go-envconfig"
)

type EnvConfig struct {
	Port int `env:"PORT,default=8081"`
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Contact the server and print out its response.
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		// defer cancel()

		name := r.URL.Query().Get("name")
		// s, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("could not greet: %v", err), http.StatusInternalServerError)
		// 	return
		// }

		// fmt.Fprintf(w, "Greeting: %s", s.GetMessage())
		fmt.Fprintf(w, "Greeting from server: %s", name)
	})

	log.Println("Starting HTTP server on", fmt.Sprintf(":%d", envConfig.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", envConfig.Port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
