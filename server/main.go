package main

import (
	"log"
	"net"

	pb "github.com/PraneGIT/gRPC-complete/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	// Serve gRPC Server on the listener
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		println("haha")
		panic(err)
	}

}
