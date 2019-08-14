package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jigadhirasu/laboratory/xgrpc/hello"
	"google.golang.org/grpc"
)

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct{}

// XHello XHello
func (*UnimplementedHelloServiceServer) XHello(ctx context.Context, req *hello.HelloTask) (*hello.HelloTask, error) {
	req.Response = "Xhello"
	return req, nil
}

//SayHello SayHello
func (*UnimplementedHelloServiceServer) SayHello(req *hello.HelloTask, srv hello.HelloService_SayHelloServer) error {
	count := 0
	for {
		count++
		err := srv.Send(&hello.HelloTask{
			Name:      fmt.Sprintf("server-%d", count),
			Timestamp: int32(time.Now().Unix()),
			Response:  "Hello",
		})
		if err != nil {
			return err
		}
		<-time.After(time.Second)
	}

}

func main() {

	log.SetFlags(log.Ltime | log.Lshortfile | log.LstdFlags)

	listener, _ := net.Listen("tcp", ":17887")
	ss := grpc.NewServer()
	hello.RegisterHelloServiceServer(ss, &UnimplementedHelloServiceServer{})

	log.Println("Start listen to tcp :17887")
	err := ss.Serve(listener)
	if err != nil {
		log.Println("NO~~~NO~~~NO~~~")
	}

}
