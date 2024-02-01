package main

import (
	"context"

	pb "github.com/PraneGIT/gRPC-complete/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	println("streaming started...")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			println("streaming ended")
			break
		}
		println(res.Message)
	}

}
