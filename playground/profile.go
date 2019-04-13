package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Server(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", Server)
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// heap
// go tool pprof http://localhost:6060/debug/pprof/heap

// cpu
// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

// trace
// wget http://localhost:6060/debug/pprof/trace?seconds=5
