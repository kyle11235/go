package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/kyle11235/go/playground/rpc/api"
)

func main() {
	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// sync call
	args := &api.Args{7, 8}
	var reply int
	err = client.Call("MyAPI.Multiply", args, &reply)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Printf("rpc %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	result := new(api.Result)
	divCall := client.Go("MyAPI.Divide", args, result, nil)
	_ = <-divCall.Done // will be equal to divCall
	fmt.Printf("rpc %d/%d=%d,%d\n", args.A, args.B, result.A, result.B)

}
