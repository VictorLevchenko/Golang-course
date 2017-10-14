package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

//Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply

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

	switch reqs[0] {
	case "GET":
		switch reqs[1] {
		case "/":
			fmt.Fprintln(con, "hangler for GET / called")
		case "/apply":
			fmt.Fprintln(con, "hangler for GET /apply called")
		default:
			fmt.Fprintln(con, "handler not found")
		}

	case "POST":
		switch reqs[1] {
		case "/":
			fmt.Fprintln(con, "hangler for POST / called")
		case "/apply":
			fmt.Fprintln(con, "hangler for POST /apply called")
		default:
			fmt.Fprintln(con, "hangler not found")
		}
	default:
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

		go serveConn(con)
	}
}
