package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)

//Take the previous program in the previous folder and change it so that:
//a template is parsed and served
//you pass data into the template

func rootHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello client!")
}

func dogHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from dog!")
}

//fname passed ass pass, lname as query parameter
func meHndl(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./routing/03/tmpl.html")
	if err != nil {
		log.Fatalf("Can't find template: %s", err)
	}
	tmpl.ExecuteTemplate(w, "tmpl.html", "Donald Trump")
}

func main() {
	log.Println("Server starting...")

	http.HandleFunc("/", rootHndl)
	http.HandleFunc("/dog/", dogHndl)
	http.HandleFunc("/me/", meHndl)

	log.Fatal("Can't serve port 8080: ", http.ListenAndServe(":8080", nil))

}
