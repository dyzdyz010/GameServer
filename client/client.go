package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var writeStr, readStr = make([]byte, 1024), make([]byte, 1024)

func main() {
	var (
		host   = "127.0.0.1"
		port   = "37576"
		remote = host + ":" + port
		reader = bufio.NewReader(os.Stdin)
	)

	con, err := net.Dial("tcp", remote)
	defer con.Close()

	if err != nil {
		fmt.Println("Server not found.")
		os.Exit(-1)
	}
	fmt.Println("Connection OK.")
	login(con)
	fmt.Println("Login succeed! Now begin to talk!")
	go read(con)

	for {
		writeStr, _, _ = reader.ReadLine()
		if string(writeStr) == "quit" {
			fmt.Println("Communication terminated.")
			os.Exit(1)
		}

		in, err := con.Write([]byte(writeStr))
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server closed. Exiting...")
			} else {
				fmt.Printf("Error when send to server: %d\n", in)
			}
			os.Exit(0)
		}

	}
}

func read(conn net.Conn) {
	for {
		length, err := conn.Read(readStr)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server closed. Exiting...")
			} else {
				fmt.Printf("Error when read from server. Error:%s\n", err)
			}
			os.Exit(0)
		}
		fmt.Println(string(readStr[:length]))
	}
}

func login(conn net.Conn) {
	// Username
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &writeStr)
	username := string(writeStr)

	// Password
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s", &writeStr)
	password := string(writeStr)

	in, err := conn.Write([]byte(username + " " + password))
	if err != nil {
		fmt.Printf("Error when send to server: %d\n", in)
		os.Exit(0)
	}

	length, err := conn.Read(readStr)
	if err != nil {
		fmt.Printf("Error when read from server. Error:%s\n", err)
		os.Exit(0)
	}
	if string(readStr[:length]) == "login error" {
		fmt.Println("Login error, try again.")
		login(conn)
	}

}
