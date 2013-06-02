package main

import (
	"GameServer/message"
	"GameServer/types"
	//"fmt"
	"net"
)

func newConn(conn net.Conn) {
	conn.Close()
}
