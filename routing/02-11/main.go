package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

//Before we WRITE our RESPONSE , let's WRITE to our RESPONSE the STATUS LINE and some
// REPONSE HEADERS. Remember the request line and status line:
//REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF
// https://tools.ietf.org/html/rfc7230#section-3.1.1
//RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP
// reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2
//Write the following strings to the response - use io.WriteString for all of the
// following except the second and third:
//"HTTP/1.1 200 OK\r\n"
//fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
//fmt.Fprint(c, "Content-Type: text/plain\r\n")
//"\r\n"
//Look in your browser "developer tools" under the network tab.
// Compare the RESPONSE HEADERS from the previous file with the
// RESPONSE HEADERS in your new solution.

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

		body := "I see you connected\n"

		io.WriteString(con, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(con, "Content-Type: text/plain\r\n")
		io.WriteString(con, "\r\n")

		io.WriteString(con, body)

		go serveConn(con)
	}
}
