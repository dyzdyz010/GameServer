package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "8000"
		remote = host + ":" + port
	)

	var (
		broadcastChan = make(chan string)
	)

	listen, err := net.Listen("tcp", remote)
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen error: ", err)
		os.Exit(-1)
	}

}
