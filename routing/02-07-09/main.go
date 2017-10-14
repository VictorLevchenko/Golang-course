package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

//Extract the code you wrote to READ from the connection using bufio.NewScanner
// into its own function called "serve".
//Pass the connection of type net.Conn as an argument into this function.
//Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

func serveConn(con net.Conn) {
	defer con.Close()
	input := bufio.NewScanner(con)
	for input.Scan() {
		txt := input.Text()
		if txt == "" {
			break
		}
		fmt.Println(txt)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Can't listen to port 8080:", err)
	}
	defer lis.Close()
	log.Println("Server starting")
	for {
		con, err := lis.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		log.Println("Connected ", con.RemoteAddr())
		io.WriteString(con, "I see you connected\n")
		go serveConn(con)
	}
}
