package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)

//Take the previous program and change it so that:
//func main uses http.Handle instead of http.HandleFunc
//Contstraint: Do not change anything outside of func main

func rootHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello client!")
}

func dogHndl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from hot dog!")
}

//fname passed ass pass, lname as query parameter
func meHndl(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./routing/05/tmpl.html")
	if err != nil {
		log.Fatalf("Can't find template: %s", err)
	}
	tmpl.ExecuteTemplate(w, "tmpl.html", "Donald Trump")
}

func main() {
	log.Println("Server starting...")

	http.Handle("/", http.HandlerFunc(rootHndl))
	http.Handle("/dog/", http.HandlerFunc(dogHndl))
	http.Handle("/me/", http.HandlerFunc(meHndl))

	log.Fatal("Can't serve port 8080: ", http.ListenAndServe(":8080", nil))

}
