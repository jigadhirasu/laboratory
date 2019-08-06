package main

// // HelloServiceServer can be embedded to have forward compatible implementations.
// type HelloServiceServer struct{}

// // SayHello SayHello
// func (*HelloServiceServer) SayHello(ctx context.Context, req *hello.HelloTask) (*hello.HelloTask, error) {
// 	req.Response = "XYZ"
// 	return req, nil
// }

// func main() {
// 	conn, _ := grpc.Dial("localhost:17887", grpc.WithInsecure())
// 	defer conn.Close()

// 	c := hello.NewHelloServiceClient(conn)
// 	res, _ := c.SayHello(context.Background(), &hello.HelloTask{
// 		Name:      "I",
// 		Timestamp: int32(time.Now().Unix()),
// 		Message:   "Hello",
// 	})

// 	log.Printf("%#v\n", res)
// }
