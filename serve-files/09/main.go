package main

import (
	"net/http"
	"fmt"
)

/*Serve the files in the "starting-files" folder
To get your images to serve, use:
func StripPrefix(prefix string, h Handler) Handler
func FileServer(root FileSystem) Handler
Constraint: you are not allowed to change the route being used for images in the template file*/

func main() {
	fmt.Println("Server starting...")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("pic"))))
	http.ListenAndServe(":8080", nil)
}
