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

type API int

func (a *API) Multiply(args *api.Args, result *int) error {
	*result = args.A * args.B
	return nil
}

func (a *API) Divide(args *api.Args, res *api.Result) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	res.A = args.A / args.B
	res.B = args.A % args.B
	return nil
}

func main() {
	api := new(API)
	rpc.Register(api)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	time.Sleep(1000 * time.Second)
}

// https://golang.org/pkg/net/rpc/