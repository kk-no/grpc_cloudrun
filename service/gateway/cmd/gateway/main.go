package main

import (
	pb "cloudrun/gen/go/proto/v1"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterCatHandlerFromEndpoint(ctx, mux, "localhost:8081", opts); err != nil {
		log.Fatalf("Failed to register groc endpoint: %v", err)
	}

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
