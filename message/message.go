package message

import (
//"fmt"
//"net"
)

func Message(broadcast, target chan string) {
	go targetedMessage(target)
}

func TargetedMessage(target chan string) {

}
