package types

import (
	//"fmt"
	"net"
)

type User struct {
	name string
	conn net.Conn
}
