// Client module
package main

import (
	"flag"
	//"fmt"
	"log"
	//"os"
)

var (
	uname  = flag.String("u", "", "Enter your username")
	passwd = flag.String("p", "", "Enter your password")
)

func main() {
	// Parse the username and password from the flags
	flag.Parse()
	if *uname == "" || *passwd == "" {
		log.Fatal("Invalid username or password: ", *uname)
	}
	// 
}
