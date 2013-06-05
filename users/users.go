/*

 文件名: users.go
 创建时间: 2013-06-02
 简介: 用户管理包，用于管理用户的数量以及相关操作。

 详情: 该包管理一个所有用户的数组，负责添加新连接的用户，删除已断开连接的用户，查找指定的用户。
 		func CheckAvailable(user *User) bool —— 检查指定用户是否已经存在在用户列表中
 		function AddUser(user User) —— 将user添加到用户数组中
 		function RemoveUser(user User) —— 把user从用户数组中删除
 		function GetUserByName(name string) User —— 根据name在用户数组中查找用户名为name的用户并返回User对象

 Copyright (C) 2013 dyzdyz010. All Rights Reserved.

*/

package users

import (
	. "GameServer/types"
	//"fmt"
)

var Users []*User

func CheckAvailable(u *User) bool {
	for _, user := range Users {
		if u.Name == user.Name {
			return false
		}
	}

	return true
}

func AddUser(u *User) {
	if CheckAvailable(u) {
		Users = append(Users, u)
	}
}

func RemoveUser(u *User) {
	for index, user := range Users {
		if u.Name == user.Name {
			Users = append(Users[:index], Users[index+1:]...)
			return
		}
	}
}

func GetUserByName(name string) *User {

	for _, user := range Users {
		if user.Name == name {
			return user
		}
	}

	return &User{}
}

func GetUserByChannel(chann chan string) {

}
