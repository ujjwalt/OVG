//server.go will provide the interface for communicating and handling hosts

// workerDead(message string), send the message and wait for ack, if no ack means worker dead

package main

import (
	"fmt"
	"github.com/ujjwalt/ovg/src/main/ovg"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path"
)

type Flag int
type Args struct {
	message string
}

func main() {

	flag := new(Flag)
	rpc.Register(flag)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil) //ListenAndServe starts an HTTP server with a given address and handler. 
	//The handler is usually nil, which means to use DefaultServeMux. 
	if err != nil {
		fmt.Println(err.Error())
	}
}

/*
TODO: Implement rpc methods to facilitate registration and message sending through the ovg package
*/
