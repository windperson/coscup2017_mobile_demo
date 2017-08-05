package main

import (
	"golang.org/x/net/context"
	pb "github.com/windperson/coscup2017-mobile-demo/go-server/proto"
	"net"
	"log"
	"google.golang.org/grpc"
)

type serverImpl struct{}

func (s *serverImpl) Greeting(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("client invoke Greeting gRPC call")
	var greeting = "Hello '" + req.Name + "' from go!"
	return &pb.HelloResponse{
		Greeting: greeting,
	}, nil
}

func (s *serverImpl) Goodbye(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("client invoke Goodbye gRPC call")
	return &pb.HelloResponse{
		Greeting: "goodbye",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &serverImpl{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}