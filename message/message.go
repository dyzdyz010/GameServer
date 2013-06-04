/*

 文件名: message.go
 创建时间: 2013-06-02
 简介: 消息中转块，包含广播与点对点通信两种分发模式。

 详情: 该包包含广播与点对点两个线程，分别负责监听对应的管道信息，并把指定信息写入对应的User.MsgChan中
 		function newConn(user types.User) —— 将user添加到用户数组中
 		function RemoveUser(user types.User) —— 把user从用户数组中删除
 		function GetUserByName(name string) types.User —— 根据name在用户数组中查找用户名为name的用户并返回types.User对象

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package message

import (
	"GameServer/types"
	"GameServer/users"

//"fmt"
//"net"
)

func Message(broadcast chan string, target chan types.TargetMsg) {
	go targetedMessage(target)

	for {
		msg := <-broadcast
		for _, user := range users.Users {
			user.MsgChan <- msg
		}
	}
}

func targetedMessage(target chan types.TargetMsg) {
	for {
		msg := <-target
		for _, name := range msg.Target {
			users.GetUserByName(name).MsgChan <- msg.Message
		}
	}
}
