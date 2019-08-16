package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/jigadhirasu/laboratory/xgrpc/hello"
	"google.golang.org/grpc"
)

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct{}

// StreamHelloStream StreamHelloStream
func (*UnimplementedHelloServiceServer) StreamHelloStream(srv hello.HelloService_StreamHelloStreamServer) error {

	go func() {
		for {
			in, err := srv.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("StoS : ", in)
		}
	}()

	tick := time.Tick(time.Second)
	for {
		ti := <-tick
		srv.Send(&hello.HelloTask{Response: ti.Format("StoS : 15:04:05")})
	}

	return nil
}

// StreamHello StreamHello
func (*UnimplementedHelloServiceServer) StreamHello(srv hello.HelloService_StreamHelloServer) error {

	for i := 0; i < 3; i++ {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println(err)
			return nil
		}

		log.Println("Sto1", in)
	}

	srv.SendAndClose(&hello.HelloTask{Response: "Sto1 Closed"})

	return nil
}

// HelloStream HelloStream
func (*UnimplementedHelloServiceServer) HelloStream(req *hello.HelloTask, srv hello.HelloService_HelloStreamServer) error {
	log.Println("1toS : ", req)
	for {
		srv.Send(&hello.HelloTask{Timestamp: int32(time.Now().Unix())})
		<-time.After(time.Second)
	}
	return nil
}

// Hello Hello
func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *hello.HelloTask) (*hello.HelloTask, error) {
	req.Response = "1to1"
	return req, nil
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
