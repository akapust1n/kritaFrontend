package main

import (
	"fmt"
	"io/ioutil"
	sw "kritaServers/backend/goserver/server"
	"net/http"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func handlerInstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("REQUEST")
	bodyBuffer, _ := ioutil.ReadAll(r.Body)

	// fmt.Printf("after parse")
	// fmt.Println(string(bodyBuffer))
	sw.InsertGeneralInfo(bodyBuffer)
}
func handlerTools(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("hadle tools!")
	bodyBuffer, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("after parse")
	fmt.Println(string(bodyBuffer))
	sw.InsertToolInfo(bodyBuffer)
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Printf("hello")
	sw.InitDB()
	defer sw.Session.Close()

	http.HandleFunc("/install/receiver/submit/org.krita.krita/", handlerInstall)
	http.HandleFunc("/tools/receiver/submit/org.krita.krita/", handlerTools)

	http.HandleFunc("/GoogleLogin", sw.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", sw.HandleGoogleCallback)

	http.ListenAndServe(":8080", nil)
}
