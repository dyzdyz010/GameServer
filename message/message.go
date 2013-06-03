package message

import (
//"fmt"
//"net"
)

func Message(broadcast, target chan string) {
	go TargetedMessage(target)
}

func TargetedMessage(target chan string) {

}
