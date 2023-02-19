package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/dollarkillerx/zim/test/grpc/proto"
	"golang.org/x/net/context"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8372")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, reply *proto.HelloReply) (*proto.HelloRequest, error) {
	fmt.Println(reply.Message)

	return &proto.HelloRequest{Name: "jxc"}, nil
}
