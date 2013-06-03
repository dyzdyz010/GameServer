package connection

import (
	//"GameServer/message"
	"GameServer/types"
	//"GameServer/users"
	//"fmt"
	"net"
)

func newConn(conn net.Conn, broadcast, target chan string) {
	conn.Close()
}

func write(user types.User) {
	for {
		msg := <-user.MsgChan
		user.Conn.Write([]byte(msg))
	}
}
