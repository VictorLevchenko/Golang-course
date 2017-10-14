package main

import (
	"net/http"
	"fmt"
	"log"
)

//ListenAndServe on port ":8080" using the default ServeMux.
//Use HandleFunc to add the following routes to the default ServeMux:
//"/" "/dog/" "/me/
//Add a func for each of the routes.
//Have the "/me/" route print out your name.


func rootHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello client!")
}

func dogHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from dog!")
}
//fname passed ass pass, lname as query parameter
func meHndl(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Donald Thrump")
}

func main() {
	log.Println("Server starting...")

	http.HandleFunc("/", rootHndl)
	http.HandleFunc("/dog/", dogHndl)
	http.HandleFunc("/me/", meHndl)

	log.Fatal("Can't serve port 8080: ", http.ListenAndServe(":8080", nil))

}
