package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jigadhirasu/laboratory/xgrpc/hello"
	"google.golang.org/grpc"
)

func main() {

	conn, _ := grpc.Dial("localhost:17887", grpc.WithInsecure())
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	stream, _ := client.SayHello(context.Background())

	go func() {
		for {
			ht, _ := stream.Recv()
			fmt.Printf("%#v \n", ht)
		}
	}()

	for i := 0; i < 10; i++ {
		stream.Send(&hello.HelloTask{Name: "I'm Hello"})
	}

	stream.CloseSend()
	<-time.After(time.Second * 3)
}
