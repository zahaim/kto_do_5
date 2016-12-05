package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var Week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

type Page struct {
	Day int
	Week []string
	Tables []*Table
}

type Table struct {
	Name string
	Value [5]bool
}

func Home(w http.ResponseWriter, req *http.Request) {
	render(w, "index.html")
}

func render(w http.ResponseWriter, tmpl string) {
	day := time.Now().Day()
	jacek := Table{Name: "Jacek", Value: [5]bool{true, false, false, false}}
	ala := Table{Name: "Ala", Value: [5]bool{false, true, true, false, false}}
	janusz := Table{Name: "Janusz", Value: [5]bool{false, true, false, true, false}}
	jan := Table{Name: "Jan", Value: [5]bool{false, false, false, false, true}}
	Tables := []*Table{&jacek,&ala,&janusz,&jan}

	values := Page{day, Week, Tables}
	tmpl = fmt.Sprintf("%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, values)
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
