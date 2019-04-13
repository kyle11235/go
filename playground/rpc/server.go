package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/kyle11235/go/playground/rpc/api"
)

type MyAPI api.API

func (s *MyAPI) Multiply(args *api.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (s *MyAPI) Divide(args *api.Args, res *api.Result) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	res.A = args.A / args.B
	res.B = args.A % args.B
	return nil
}

func main() {
	myAPI := new(MyAPI)
	rpc.Register(myAPI)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	time.Sleep(1000 * time.Second)
}

// https://golang.org/pkg/net/rpc/