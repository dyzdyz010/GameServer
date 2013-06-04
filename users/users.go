/*

 文件名: users.go
 创建时间: 2013-06-02
 简介: 用户管理包，用于管理用户的数量以及相关操作。

 详情: 该包管理一个所有用户的数组，负责添加新连接的用户，删除已断开连接的用户，查找指定的用户。
 		function AddUser(user types.User) —— 将user添加到用户数组中
 		function RemoveUser(user types.User) —— 把user从用户数组中删除
 		function GetUserByName(name string) types.User —— 根据name在用户数组中查找用户名为name的用户并返回types.User对象

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package users

import (
	"GameServer/types"
	//"fmt"
)

var Users = make([]types.User, 10000)

func AddUser(user types.User) {

}

func RemoveUser(user types.User) {

}

func GetUserByName(name string) types.User {

	for _, user := range Users {
		if user.Name == name {
			return user
		}
	}

	return types.User{}
}

func GetUserByChannel(chann chan string) {

}
