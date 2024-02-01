package main

import (
	"io"
	"log"

	pb "github.com/PraneGIT/gRPC-complete/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name:%v", msg.Name)
		messages = append(messages, "hello "+msg.Name)
	}
}
