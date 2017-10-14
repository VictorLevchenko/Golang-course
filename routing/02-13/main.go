package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
)

//Print to standard out (the terminal) the REQUEST method and the REQUEST URI from
// the REQUEST LINE.
//Add this data to your REPONSE so that this data is displayed in the browser.

func serveConn(con net.Conn) {
	defer con.Close()
	input := bufio.NewScanner(con)
	i := 0;
	req := ""
	for input.Scan() {
		txt := input.Text()
		if txt == "" {
			break
		}
		if i == 0 {
			req = txt
		}
		i++
		fmt.Println(txt)
	}
	reqs := strings.Split(req, " ")

	body := fmt.Sprintf("I see you connected  to the uri %s with method %s\n ", reqs[1], reqs[0])

	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/plain\r\n")
	io.WriteString(con, "\r\n")

	io.WriteString(con, body)
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

		go serveConn(con)
	}
}
