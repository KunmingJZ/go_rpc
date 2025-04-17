package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mirco_goods/proto/greeterService"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type AddGoodsRequest struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsResponse struct {
	Success bool
	Message string
}

func RpcPattern() {
	fmt.Println("Hello, world! client")
	conn, err := rpc.Dial("tcp", "localhost:8020")
	if err != nil {
		fmt.Println("dialing error:", err)
		return
	}
	defer conn.Close()

	var result AddGoodsResponse

	request := AddGoodsRequest{
		Id:      1,
		Title:   "hello",
		Price:   10.0,
		Content: "world<><>< from client",
	}
	err2 := conn.Call("Goods.AddGoods", request, &result)
	if err2 != nil {
		fmt.Println("rpc error:", err2)
		return
	}
	fmt.Printf("result:%#v\n", result)
}

func JsonPattern() {
	fmt.Println("Hello, world! json client")
	conn, err := net.Dial("tcp", "localhost:8020")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call("HelloRequest.SayHello", "hello ,.,!from client", &reply)
	if err != nil {
		fmt.Println("dialing error:", err)
		panic(err)
	}
	fmt.Println("client received:", reply)

}

func RpcGreeder() {
	grpcClient, err := grpc.Dial("localhost:8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc dial error:", err)
	}
	client := greeterService.NewGreeterClient(grpcClient)
	//res := &greeterService.GreetingResponse{}
	var res, _ = client.SayHello(context.Background(), &greeterService.GreetingRequest{Name: "hello from client", Language: "en"})
	fmt.Printf("grpc client received: %#v", res)
}

func main() {
	//JsonPattern()
	RpcGreeder()
}
