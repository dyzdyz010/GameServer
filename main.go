package main

import (
	"GameServer/connection"
	"GameServer/console"
	"GameServer/message"

	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "37576"
		remote = host + ":" + port
	)

	// 初始化
	initialize()

	// 监听端口
	listen, err := net.Listen("tcp", remote)
	defer listen.Close()
	if err != nil {
		fmt.Println("Listen error: ", err)
		os.Exit(-1)
	}
	fmt.Println("Server start listen on " + remote + "\n")

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

func initialize() {
	fmt.Println("\nServer initializing...\n")

	// 开启多核
	runtime.GOMAXPROCS(4)
	fmt.Println("Multi-CPU support active, current CPUs in use: ", runtime.NumCPU(), "\n")

	// 开启控制台
	go console.Console()
	fmt.Println("Server console activated.\n")

	// 准备通信线程
	go message.Message()
	fmt.Println("Message routines ready...\n")
}
