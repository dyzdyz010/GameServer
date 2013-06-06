package console

import (
	"GameServer/users"

	"bufio"
	"fmt"
	"os"
	"runtime"
	"syscall"
)

var (
	cpus     = "cpu num"
	routines = "goroutines"
	usernum  = "user num"
	userlist = "user list"
	memusage = "memory usage"
	quit     = "quit"
)

var (
	command = make([]byte, 1024)
	reader  = bufio.NewReader(os.Stdin)
)

func Console() {
	for {
		command, _, _ = reader.ReadLine()
		switch string(command) {

		case quit:
			fmt.Println("Server stopped.")
			os.Exit(0)

		case cpus:
			fmt.Println("\nThe number of CPUs currently in use: ", runtime.NumCPU())

		case routines:
			fmt.Println("\nCurrent number of goroutines: ", runtime.NumGoroutine())

		case usernum:
			fmt.Println("The number of clients currently online is ", users.NumberOfUsers())

		case userlist:
			fmt.Println("Currently online users: ")
			fmt.Println(users.Users)

		case memusage:
			us := &syscall.Rusage{}
			err := syscall.Getrusage(syscall.RUSAGE_SELF, us)
			if err != nil {
				fmt.Println("Get usage error: ", err, "\n")
			} else {
				fmt.Printf("\nMemory Usage: %f MB\n\n", float64(us.Maxrss)/1024/1024)
			}

		default:
			fmt.Println("Command error, try again.")
		}

	}
}
