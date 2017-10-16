package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

//Serve the files in the "starting-files" folder
//To get your images to serve, use only this:
//fs := http.FileServer(http.Dir("public"))
//Hint: look to see what type FileServer returns, then think it through.
func indexHndl(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal("Template not found", err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)

}

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", indexHndl)
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}
