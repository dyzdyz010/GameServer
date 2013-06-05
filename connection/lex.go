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
