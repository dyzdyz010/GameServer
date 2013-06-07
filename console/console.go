/*

 文件名: console.go
 创建时间: 2013-06-06
 简介: 给服务端添加一个控制台，便于输入指令调试。

 详情: 该包包含一个在服务端等待I/O操作的函数，用来对输入的指令进行回应。
 		function Console() —— 开启控制台routine并负责功能

 		指令：
 			cpu num ——　显示程序能使用的ＣＰＵ数量
 			goroutines —— 查看当前程序的routine数量
 			user num —— 查看当前在线用户数量
 			user list —— 查看当前在线用户列表
 			memory usage —— 查看程序当前的内存使用情况
 			quit —— 关闭服务器


 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package console

import (
	"GameServer/users"

	"bufio"
	"fmt"
	"os"
	"runtime"
	"syscall"
)

var (
	cpus     = "cpu num"
	routines = "goroutines"
	usernum  = "user num"
	userlist = "user list"
	memusage = "memory usage"
	quit     = "quit"
)

var (
	command = make([]byte, 1024)
	reader  = bufio.NewReader(os.Stdin)
)

func Console() {
	for {
		command, _, _ = reader.ReadLine()
		switch string(command) {

		case quit:
			fmt.Println("Server stopped.")
			os.Exit(0)

		case cpus:
			fmt.Println("\nThe number of CPUs currently in use: ", runtime.NumCPU())

		case routines:
			fmt.Println("\nCurrent number of goroutines: ", runtime.NumGoroutine())

		case usernum:
			fmt.Println("The number of clients currently online is ", users.NumberOfUsers())

		case userlist:
			fmt.Println("Currently online users: ")
			fmt.Println(users.Users)

		case memusage:
			us := &syscall.Rusage{}
			err := syscall.Getrusage(syscall.RUSAGE_SELF, us)
			if err != nil {
				fmt.Println("Get usage error: ", err, "\n")
			} else {
				fmt.Printf("\nMemory Usage: %f MB\n\n", float64(us.Maxrss)/1024/1024)
			}

		default:
			fmt.Println("Command error, try again.")
		}

	}
}
