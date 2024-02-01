package main

import (
	pb "github.com/PraneGIT/gRPC-complete/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"praneki", "pranjal", "jha"},
	}

	// Call the remote method
	// callSayHello(client)

	// callSayHelloServerStreaming(client, names)

	callSayHelloClientStream(client, names)
}
