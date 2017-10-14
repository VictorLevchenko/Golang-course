package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

//In that previous exercise, we WROTE to the connection.
//Now I want you to READ from the connection.
//You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.
//Use bufio.NewScanner() to read from the connection.
//After all of the reading, include these lines of code:
//fmt.Println("Code got here.") io.WriteString(c, "I see you connected.")
//Launch your TCP server.
//In your web browser, visit localhost:8080.
//Now go back and look at your terminal.
//Can you answer the question as to why "I see you connected." is never written?
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

		input := bufio.NewScanner(con)
		input.Scan()
		txt := input.Text()
		fmt.Println("first token scanned", txt)
		fmt.Println("Code got here")

		io.WriteString(con, "I see you connected\n")

		con.Close()
	}
}
