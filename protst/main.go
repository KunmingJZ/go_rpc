package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"mirco_goods/proto/greeterService"
	"mirco_goods/proto/userService"
	"net"
)

func serialize(msg proto.Message) {
	fmt.Println("Hello, world!")
	u := &userService.UserInfo{
		Age:     1,
		Name:    "Alice",
		Email:   "alice@example.com",
		Hobbies: []string{"reading", "swimming", "running"},
	}
	var name = u.GetName()
	fmt.Println(u)
	fmt.Println(name)
	data, _ := proto.Marshal(u)
	fmt.Println(data)
	user := &userService.UserInfo{}
	err := proto.Unmarshal(data, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", user)
}

type Hello struct{}

func (h Hello) SayHello(c context.Context, request *greeterService.GreetingRequest) (*greeterService.GreetingResponse, error) {
	fmt.Println("-----say Hello, world---")
	return &greeterService.GreetingResponse{
		Message: "Hello, " + request.Name + "!",
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	greeterService.RegisterGreeterServer(grpcServer, &Hello{})
	listener, err := net.Listen("tcp", ":8020")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	grpcServer.Serve(listener)
}
