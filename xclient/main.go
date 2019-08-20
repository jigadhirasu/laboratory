package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/jigadhirasu/laboratory/xgrpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	log.SetFlags(log.Llongfile)

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, _ := grpc.Dial(":17887", grpc.WithTransportCredentials(creds))
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	// 1to1
	log.Println(client.Hello(context.Background(), &hello.Empty{}))

	// Sto1
	stream1, err := client.StreamHello(context.Background())
	if err != nil {
		log.Println(err)
	}

	stream1.Send(&hello.HelloTask{})
	stream1.Send(&hello.HelloTask{})
	stream1.Send(&hello.HelloTask{})
	h1 := &hello.HelloTask{}
	stream1.RecvMsg(h1)
	log.Println(h1)

	// 1toS

	stream2, err := client.HelloStream(context.Background(), &hello.HelloTask{Name: "ABC"})
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < 3; i++ {
		log.Println(stream2.Recv())
	}
	stream2.CloseSend()

	// StoS

	stream3, err := client.StreamHelloStream(context.Background())
	for i := 0; i < 3; i++ {
		stream3.Send(&hello.HelloTask{Name: fmt.Sprintf("ABC%d", i+1)})

		ht, err := stream3.Recv()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(ht)
	}
	stream3.CloseSend()

}
