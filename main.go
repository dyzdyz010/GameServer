/*

 文件名: main.go
 创建时间: 2013-06-02
 简介: 服务器主函数

 详情: 程序入口，负责初始化CPU分配、打开控制台、准备消息分发routines，侦听端口，并且接收客户端的连接。

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

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
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Multi-CPU support active, current CPUs in use: ", runtime.NumCPU(), "\n")

	// 开启控制台
	go console.Console()
	fmt.Println("Server console activated.\n")

	// 准备通信线程
	go message.Message()
	fmt.Println("Message routines ready...\n")
}
