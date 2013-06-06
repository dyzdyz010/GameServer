/*

 文件名: types.go
 创建时间: 2013-06-02
 简介: 数据类型管理包，管理整个程序中的所有自定义类型，如用户对象等。

 详情: 该包中所包含的数据结构：
 		struct User —— 用户对象，包括名称、密码、通信通道、连接句柄等四个成员
 		struct BroadcastMsg　——　广播通信包对象，包含来源、要传递的信息等两个成员
 		struct TargetMsg —— 点对点通信包对象，包含来源、目标用户数组、要传递的信息等三个成员

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package types

import (
	"net"
)

type User struct {
	Name     string
	Password string
	MsgChan  chan string
	Conn     net.Conn
}

type BroadcastMsg struct {
	Origin  string
	Message string
}

type TargetMsg struct {
	Origin  string
	Target  []string
	Message string
}

var (
	BroadcastChan = make(chan BroadcastMsg)
	TargetMsgChan = make(chan TargetMsg)
)

func (u *User) String() string {
	return "[Name: " + u.Name + ", Password: " + u.Password + ", Address: " + u.Conn.RemoteAddr().String() + "]"
}
