/*

 文件名: lex.go
 创建时间: 2013-06-03
 简介: 客户端消息解析函数。

 详情: 该文件仅有一个函数：
 		function lex(string, *User)
 		将客户端发来的消息进行解析，确定分发对象，并把消息包结构体发送给message包。

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package connection

import (
	. "GameServer/types"
	"fmt"
	"strings"
)

func lex(commingMsg string, u *User) {
	arr1 := strings.Split(commingMsg, "$to ")
	fmt.Println("arr1: ", arr1)
	arr2 := strings.Split(arr1[len(arr1)-1], "$say ")
	fmt.Println("arr2: ", arr2)
	msg := arr2[1]

	if arr2[0] == "" {
		fmt.Println("Broadcast: ", msg)
		bag := BroadcastMsg{u.Name, msg}
		BroadcastChan <- bag
	} else {
		targets := strings.Split(arr2[0], ",")
		fmt.Println("To: ", targets)
		trim(targets)
		bag := TargetMsg{u.Name, targets, msg}
		TargetMsgChan <- bag
	}
}

func trim(arr []string) []string {
	for i, str := range arr {
		str = strings.Trim(str, " ")
		arr[i] = str
	}

	return arr
}
