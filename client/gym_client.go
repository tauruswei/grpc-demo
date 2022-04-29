package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/lightweight"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGymClient(conn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)

	defer cancelFunc()

	r, err := c.BodyBuilding(ctx, &pb.Person{
		Name:    "daxiongdi",
		Actions: []string{"卧推", "硬拉", "深蹲"},
	})

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("code: %d, msg: %s", r.Code, r.Msg)
}
