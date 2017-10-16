package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

//ListenAndServe on port 8080 of localhost
//For the default route "/" Have a func called "foo" which writes to the response "foo ran"
//For the route "/dog/" Have a func called "dog" which parses a template called
// "dog.gohtml" and writes to the response "
//This is from dog
//" and also shows a picture of a dog when the template is executed.
//Use "http.ServeFile" to serve the file "dog.jpeg"

func foo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal("template is missing:", err)
	}
	tmpl.ExecuteTemplate(w, "dog.gohtml", "This is from dog")
}

func dogPic(w http.ResponseWriter, r * http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/dog.jpeg", dogPic )
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}
