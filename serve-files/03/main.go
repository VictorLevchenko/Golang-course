package main

import (
	"net/http"

	"fmt"
)

//Serve the files in the "starting-files" folder
//Use "http.FileServer"

func main() {
	fmt.Println("Server starting...")
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
