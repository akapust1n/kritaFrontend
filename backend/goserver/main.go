package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	sw "kritaServers/backend/goserver/server"
	"net/http"
	"time"
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
	//fmt.Println(string(bodyBuffer))
	sw.InsertGeneralInfo(bodyBuffer)
}
func handlerTools(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(bodyBuffer))
	sw.InsertToolInfo(bodyBuffer)
}
func handlerImageProperties(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(bodyBuffer))
	sw.InsertImageInfo(bodyBuffer)
}
func handlerAsserts(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Asserts")
	fmt.Println(string(bodyBuffer))
	//	sw.InsertAssertInfo(bodyBuffer)
}
func handlerAgregatedData(w http.ResponseWriter, r *http.Request) {
	dataType := r.URL.Query().Get("datatype")
	result, err := json.Marshal(sw.GetAgregatedData(dataType))

	sw.CheckErr(err)
	w.Write(result)
}
func handlerActions(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(bodyBuffer))
	sw.InsertActionInfo(bodyBuffer)
}

func main() {
	fmt.Printf("hello")
	sw.InitDB()
	defer sw.Session.Close()

	http.HandleFunc("/install/receiver/submit/org.krita.krita/", handlerInstall)
	http.HandleFunc("/tools/receiver/submit/org.krita.krita/", handlerTools)
	http.HandleFunc("/imageProperties/receiver/submit/org.krita.krita/", handlerImageProperties)
	http.HandleFunc("/asserts/receiver/submit/org.krita.krita/", handlerAsserts)
	http.HandleFunc("/actions/receiver/submit/org.krita.krita/", handlerActions)

	http.HandleFunc("/agregatedData", handlerAgregatedData)

	http.HandleFunc("/GoogleLogin", sw.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", sw.HandleGoogleCallback)
	ticker := time.NewTicker(time.Hour * 1)

	go func() {
		for t := range ticker.C {
			sw.AgregateInstalInfo()
			fmt.Println("Tick at", t)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
