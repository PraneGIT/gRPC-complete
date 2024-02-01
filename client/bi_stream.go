package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/PraneGIT/gRPC-complete/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			log.Printf("got response: %v", res.Message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("bidirectional streaming finished")

}
