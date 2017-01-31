package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

type Clicks struct {
  First string
  Second string
}

func asknread(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login2.gtpl")
        values := Clicks{"checked",""}
        t.Execute(w, values)
    } else {
        r.ParseForm()
        fmt.Println("first:", r.Form["first"])
        fmt.Println("second:", r.Form["second"])
        fmt.Println("whole Form:", r.Form)
        values := Clicks{"","checked"}
        t, _ := template.ParseFiles("login2.gtpl")
        t.Execute(w, values)
    }
}

func main() {
    http.HandleFunc("/", asknread)
    err := http.ListenAndServe(":8080 ", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
