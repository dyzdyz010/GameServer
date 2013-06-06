/*

 文件名: connection.go
 创建时间: 2013-06-02
 简介: 用户建立与客户端的连接，进行读写数据。

 详情: 该包包含一个与客户端的连接所需的读与写线程，分别对客户端的信息进行读写。
 		function NewConnection(net.Conn) —— 创建新的与客户端的连接线程，并负责读取客户端信息
 		function write(*User) —— 独立于读线程的写线程，监听User.MsgChan管道中的消息并写回客户端
 		function disconnect(*User) —— 断开与客户端的连接，销毁相关资源
 		function login(*User) —— 负责处理用户登录信息验证

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package connection

import (
	. "GameServer/types"
	"GameServer/users"

	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
)

var (
	commingData = make([]byte, 1024)
)

func NewConnection(conn net.Conn) {
	fmt.Printf("New connection %s created.\n", conn.RemoteAddr().String())
	defer conn.Close()

	// 创建用户结构
	msgChan := make(chan string)
	user := &User{"", "", msgChan, conn}

	// 开启写线程
	go write(user)

	// 开始读入客户端信息
	login(user)
	BroadcastChan <- BroadcastMsg{user.Name, "entered the game."}
	for {
		lengh, err := conn.Read(commingData)

		// 如果读入数据时出错，通知写回线程退出，广播退出信息，关闭用户连接
		if err != nil {
			if err == io.EOF {
				fmt.Println(user.Name + " has left the game.")
			} else {
				fmt.Println("Read error: ", err)
			}
			disconnect(user)
			return
		}

		commingStr := string(commingData[:lengh])
		lex(commingStr, user)
	}
}

func write(u *User) {
	for {
		msg := <-u.MsgChan
		if msg == "close" {
			runtime.Goexit()
		}
		u.Conn.Write([]byte(msg))
	}
}

func disconnect(u *User) {
	u.MsgChan <- "close"
	BroadcastChan <- BroadcastMsg{u.Name, "has left the game."}
	users.RemoveUser(u)
	runtime.Goexit()
}

func login(u *User) {
	// 获取登录信息
	length, err := u.Conn.Read(commingData)
	if err != nil {
		fmt.Println("Login error: ", err)
		disconnect(u)
	}

	loginArr := strings.Split(string(commingData[:length]), " ")
	if users.CheckAvailable(u) {
		u.Name = loginArr[0]
		u.Password = loginArr[1]
		users.AddUser(u)
	} else {
		u.MsgChan <- "login error"
		login(u)
	}
}
