package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	render(w, "index.html")
}

func render(w http.ResponseWriter, tmpl string) {
	tmpl = fmt.Sprintf("%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
