package main

import (
	"log"
	"net/http"
	"strconv"
	"html/template"
)
var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Can't find template. ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var count int
	cookie, err := r.Cookie("visited")
	if err == nil {
		count, _ = strconv.Atoi(cookie.Value)
	}
	count ++
	strCount := strconv.Itoa(count)

	http.SetCookie(w, &http.Cookie{
		Name:  "visited",
		Value: strCount,
	})
	tmpl.ExecuteTemplate(w, "index.gohtml", count)
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
