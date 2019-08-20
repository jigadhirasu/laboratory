package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
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
func (*UnimplementedHelloServiceServer) Hello(ctx context.Context, req *hello.Empty) (*hello.HelloTask, error) {
	return &hello.HelloTask{Response: "1to1"}, nil
}

func main() {

	log.SetFlags(log.Ltime | log.Lshortfile | log.LstdFlags)

	address := ""
	port := "17887"

	cer, err := tls.LoadX509KeyPair("xx.crt", "xx.key")
	if err != nil {
		log.Println(err)
		log.Println("Load x509 keyPair Error")
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	lis, err := tls.Listen("tcp", address+":"+port, config)
	if err != nil {
		log.Println(err)
		log.Println("listen by TLS Error")
		return
	}

	ss := grpc.NewServer()
	hello.RegisterHelloServiceServer(ss, &UnimplementedHelloServiceServer{})

	log.Println("Start listen to tcp " + address + ":" + port)
	err = ss.Serve(lis)
	if err != nil {
		log.Println("NO~~~NO~~~NO~~~")
		return
	}

}
