package main

import (
	"GameServer/connection"
	"GameServer/message"

	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "37576"
		remote = host + ":" + port
	)

	// 监听端口
	listen, err := net.Listen("tcp", remote)
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen error: ", err)
		os.Exit(-1)
	}

	// 准备通信线程
	go message.Message()

	// 等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Accept error: ", err)
			continue
		}
		go connection.NewConnection(conn)
	}

}
