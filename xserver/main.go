package main

import (
	"github.com/jigadhirasu/laboratory/xgrpc/hello"
)

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct{}

//SayHello SayHello
func (*UnimplementedHelloServiceServer) SayHello(srv hello.HelloService_SayHelloServer) error {
	// for {
	// 	in, err := srv.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}

	// }
	return nil
}

func main() {

}
