package main

import (
	"io"
	"log"
	"net"

	"github.com/juliusv/grpc-streaming-test/hello"
	"google.golang.org/grpc"
)

type server struct{}

func (_ *server) Hello(stream hello.Hello_HelloServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			log.Println("Received EOF")
			return nil
		}
		if err != nil {
			log.Println("Received error:", err)
			return err
		}
		log.Println("Received hello request.")
	}
}

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
