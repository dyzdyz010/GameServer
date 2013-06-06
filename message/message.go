/*

 文件名: message.go
 创建时间: 2013-06-02
 简介: 消息中转块，包含广播与点对点通信两种分发模式。

 详情: 该包包含广播与点对点两个线程，分别负责监听对应的管道信息，并把指定信息写入对应的User.MsgChan中
 		function Message() —— 包入口，负责开辟点对点消息分发线程，并监听广播消息并分发
 		function targetedMessage() —— 点对点消息监听及分发函数，负责监听点对点消息并分发给目标用户

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package message

import (
	. "GameServer/types"
	"GameServer/users"

	//"fmt"
	//"net"
)

func Message() {
	go targetedMessage()

	for {
		bag := <-BroadcastChan
		go func(bag BroadcastMsg) {
			for _, user := range users.Users {
				msg := bag.Origin + ": " + bag.Message
				user.MsgChan <- msg
			}
		}(bag)
	}
}

func targetedMessage() {
	for {
		bag := <-TargetMsgChan
		for _, name := range bag.Target {
			msg := bag.Origin + ": " + bag.Message
			users.GetUserByName(name).MsgChan <- msg
		}
	}
}
