package main

import (
	"encoding/json"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("REQUEST")
	bodyBuffer, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("after parse")
	fmt.Println(string(bodyBuffer))
	sw.InsertGeneralInfo(bodyBuffer)
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	fmt.Printf("hello")
	var m map[string]interface{}
	sw.InitDB()
	defer sw.Session.Close()

	jsonString := `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return
	}

	c := sw.Session.DB("d").C("collectio2n")
	c.Insert(m)

	http.HandleFunc("/receiver/submit/org.krita.krita", handler)
	http.HandleFunc("/GoogleLogin", sw.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", sw.HandleGoogleCallback)

	http.ListenAndServe(":8080", nil)
}
