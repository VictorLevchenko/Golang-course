package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

//Serve the files in the "starting-files" folder
//To get your images to serve, use:
//func StripPrefix(prefix string, h Handler) Handler
//func FileServer(root FileSystem) Handler
//Constraint: you are not allowed to change the route being used for images in the template file

func indexHndl(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalf("Can't find template. %s", err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	fmt.Println("Server starting...")
	http.Handle("/resources/pics", http.StripPrefix("/resources/pics",
		http.FileServer(http.Dir("public/pics"))))
	http.HandleFunc("/", indexHndl)
	http.ListenAndServe(":8080", nil)
}
