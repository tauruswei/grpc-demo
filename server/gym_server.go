package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/lightweight"
	"log"
	"net"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnimplementedGymServer
}

func (s *Server) BodyBuilding(ctx context.Context, person *pb.Person) (*pb.Reply, error) {
	fmt.Printf("%s 正在健身，动作：%s \n", person.Name, person.Actions)
	return &pb.Reply{Code: 0, Msg: "ok"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGymServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
