package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type Goods struct{}

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

func (g Goods) AddGoods(req AddGoodsRequest, res *AddGoodsResponse) error {
	fmt.Println("AddGoods called...", req)
	*res = AddGoodsResponse{Success: true, Message: "Goods added successfully"}
	return nil
}

type HelloRequest struct {
	Hello string
}

func (hello *HelloRequest) SayHello(req string, res *string) error {
	*res = "Hello =jon=" + req
	fmt.Println("SayHello called...", req)
	return nil
}

func main() {
	fmt.Println("Hello, world! Goods service started.")
	error := rpc.RegisterName("Goods", new(Goods))
	if error != nil {
		fmt.Println("Error registering Goods service:", error)
	}
	error1 := rpc.RegisterName("HelloRequest", new(HelloRequest))
	if error1 != nil {
		fmt.Println("Error registering Hello service:", error1)
	}
	listener, error := net.Listen("tcp", "127.0.0.1:8020")
	if error != nil {
		fmt.Println("Error listening:", error)
	}
	defer listener.Close()

	for {
		fmt.Println("Goods service listening on 127.0.0.1:8020")
		conn, error := listener.Accept()
		if error != nil {
			fmt.Println("Error accepting:", error)
			continue
		}
		// go rpc.ServeConn(conn
		// json.NewServerConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
