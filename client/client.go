package main

import (
	"log"
	"time"

	"github.com/juliusv/grpc-streaming-test/hello"
	"google.golang.org/grpc"

	"golang.org/x/net/context"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	stream, err := hello.NewHelloClient(conn).Hello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for {
		err := stream.Send(&hello.HelloRequest{})
		if err != nil {
			log.Println("Error sending:", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
