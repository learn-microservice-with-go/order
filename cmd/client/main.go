package main

import (
	"context"
	"log"

	proto "github.com/learn-frame/learn-micro-service/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//获得一个 Client 连接
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// 获得一个 ChatService 的 client
	c := proto.NewUserServiceClient(conn)

	// grpc 调用远程的 GetUser
	response, err := c.GetUser(context.Background(), &proto.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
