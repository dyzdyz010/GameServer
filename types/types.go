/*

 文件名: types.go
 创建时间: 2013-06-02
 简介: 数据类型管理包，管理整个程序中的所有自定义类型，如用户对象等。

 详情: 该包中所包含的数据结构：
 		struct User —— 用户对象，包括名称、密码、通信通道、连接句柄等四个成员
 		struct TargetMsg —— 点对点通信包对象，包含目标用户数组、要传递的信息等两个成员

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package types

import (
	//"fmt"
	"net"
)

type User struct {
	Name     string
	Password string
	MsgChan  chan string
	Conn     net.Conn
}

type TargetMsg struct {
	Target  []string
	Message string
}
