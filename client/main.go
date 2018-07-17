package main

import (
	"fmt"
	"log"

	pb "github.com/delznet/snowflake/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7890", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewSnowflakeClient(conn)
	fmt.Println(c.Gen(context.Background(), &pb.SnowflakeRequest{Format:"int64"}))

}
