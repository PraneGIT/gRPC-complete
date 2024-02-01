package main

import (
	"context"
	"log"
	"time"

	pb "github.com/PraneGIT/gRPC-complete/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("client streaming")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		panic(err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			panic(err)
		}

		log.Printf("send req withname:%s", name)
		time.Sleep((time.Second * 2))
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming ended")

	if err != nil {
		panic(err)
	}
	log.Printf("\nGot HelloResponse %v", res)

}
