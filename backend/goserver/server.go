package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)
var data string

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "KRITA REQUEST ")
    fmt.Printf("REQUEST")
    bodyBuffer, _ := ioutil.ReadAll(r.Body)  
    data = data+ string(bodyBuffer)
    fmt.Printf(string(bodyBuffer))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Last requests</h1><div>%s</div>", data)
}

func main() {
    http.HandleFunc("/receiver/submit/org.krita.krita", handler)
    http.HandleFunc("/", viewHandler)

    http.ListenAndServe(":8080", nil)
}
