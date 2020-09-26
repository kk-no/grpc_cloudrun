package main

import (
	pb "cloudrun/gen/go/proto/v1"
	"cloudrun/service/cat"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}

	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	pb.RegisterCatServer(server, &cat.MyCatService{})
	if err := server.Serve(listenPort); err != nil {
		log.Fatal("serve end")
	}
}
