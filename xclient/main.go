package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jigadhirasu/laboratory/xgrpc/hello"
	"google.golang.org/grpc"
)

func main() {

	conn, _ := grpc.Dial(":17888", grpc.WithInsecure())
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	stream, _ := client.SayHello(context.Background(), &hello.HelloTask{})

	times := 0
	for {
		ht, _ := stream.Recv()
		fmt.Printf("%#v \n", ht)
		times++
		if times > 10 {
			break
		}
		<-time.After(time.Second)
	}
	stream.CloseSend()
}
