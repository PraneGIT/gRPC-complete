package main

import (
	"context"
	"time"

	pb "github.com/PraneGIT/gRPC-complete/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the remote method
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		panic(err)
	}

	println(res.Message)
}
