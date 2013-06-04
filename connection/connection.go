/*

 文件名: connection.go
 创建时间: 2013-06-02
 简介: 用户建立与客户端的连接，进行读写数据。

 详情: 该包包含一个与客户端的连接所需的读与写线程，分别对客户端的信息进行读写。
 		function NewConnection(net.Conn, chan string, chan types.TargetMsg) —— 创建新的与客户端的连接线程
 		function write(types.User) —— 独立于读线程的写线程，监听User.MsgChan管道中的消息并写回客户端

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package connection

import (
	"GameServer/types"
	//"GameServer/users"
	"fmt"
	"net"
)

var (
	commingData = make([]byte, 1024)
)

func NewConnection(conn net.Conn, broadcast chan string, target chan types.TargetMsg) {
	fmt.Printf("New connection %s created.", conn.RemoteAddr().String())
	msgChan := make(chan string)
	user := types.User{"", "", msgChan, conn}
	go write(user)
	for {
		lengh, err := conn.Read(commingData)

		// 如果读入数据时出错，通知写回线程退出，广播退出信息，关闭用户连接
		if err != nil {
			fmt.Println("Read error: ", err)
			msgChan <- "close"
			conn.Close()
			return
		}

		commingStr := string(commingData[:lengh])
		fmt.Println(commingStr)
	}
}

func write(user types.User) {
	for {
		msg := <-user.MsgChan
		if msg == "close" {
			return
		}
		user.Conn.Write([]byte(msg))
	}
}
