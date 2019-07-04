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
	var result1 int
	err = client.Call("API.Multiply", args, &result1)
	if err != nil {
		log.Fatal("rpc error:", err)
	}
	fmt.Printf("rpc %d*%d=%d\n", args.A, args.B, result1)

	// Asynchronous call
	result2 := new(api.Result)
	divCall := client.Go("API.Divide", args, result2, nil)
	_ = <-divCall.Done // will be equal to divCall
	fmt.Printf("rpc %d/%d=%d,%d\n", args.A, args.B, result2.A, result2.B)

}
