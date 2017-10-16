package main

import (
	"net/http"
	"log"
)

/*
Starting with the code in the "starting-files" folder:
wire this program up so that it works
ParseGlob in an init function
Use HandleFunc for each of the routes
Combine apply & applyProcess into one func called "apply"
Inside the func "apply", use this code to create the logic to respond differently
 to a POST method and a GET method
if req.Method == http.MethodPost {
// code here
return
}
func StripPrefix(prefix string, h Handler) Handler
func FileServer(root FileSystem) Handler
*/

func main() {
//TODO not solved yet
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
