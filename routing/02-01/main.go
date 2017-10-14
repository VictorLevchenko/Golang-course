package main

import (
	"net"
	"log"
	"io"
)

//Create a basic server using TCP
//The server should use net.Listen to listen on port 8080.
//Remember to close the listener using defer.
//Remember that from the "net" package you first need to LISTEN, then
// you need to ACCEPT an incoming connection.
//Now write a response back on the connection.
//Use io.WriteString to write the response: I see you connected.
//Remember to close the connection.
//Once you have all of that working, run your TCP server and test it from
// telnet (telnet localhost 8080).
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
		io.WriteString(con, "I see you connected\n")
		log.Println("Connected ", con.RemoteAddr())
		con.Close()
	}
}
