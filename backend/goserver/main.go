package main

import (
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
	fmt.Println(string(bodyBuffer))
	sw.InsertGeneralInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func handlerTools(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))
	sw.InsertToolInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func handlerImageProperties(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))
	sw.InsertImageInfo(bodyBuffer)
}
func handlerAsserts(w http.ResponseWriter, r *http.Request) {
	//	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Asserts")
	//	fmt.Println(string(bodyBuffer))
	//	sw.InsertAssertInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
func handlerActions(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))
	sw.InsertActionInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
func handlerGetTools(w http.ResponseWriter, r *http.Request) {
	temp := sw.Agregated("tools")
	fmt.Fprintf(w, temp)
}
func handlerGetActions(w http.ResponseWriter, r *http.Request) {
	temp := sw.Agregated("actions")
	fmt.Fprintf(w, temp)
}
func handlerGetInstallInfo(w http.ResponseWriter, r *http.Request) {
	temp := sw.Agregated("install")
	fmt.Println("HANDLE INSTALL GET")
	fmt.Fprintf(w, temp)
}
func handlerGetImageInfo(w http.ResponseWriter, r *http.Request) {
	temp := sw.Agregated("images")
	fmt.Fprintf(w, temp)
}

func main() {
	fmt.Printf("hello")
	sw.InitDB()
	defer sw.Session.Close()

	http.HandleFunc("/install/receiver/submit/org.krita.krita", handlerInstall)
	http.HandleFunc("/tools/receiver/submit/org.krita.krita", handlerTools)

	http.HandleFunc("/imageProperties/receiver/submit/org.krita.krita", handlerImageProperties)
	http.HandleFunc("/asserts/receiver/submit/org.krita.krita", handlerAsserts)
	http.HandleFunc("/actions/receiver/submit/org.krita.krita", handlerActions)

	http.HandleFunc("/GoogleLogin", sw.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", sw.HandleGoogleCallback)

	http.HandleFunc("/get/tools", handlerGetTools)
	http.HandleFunc("/get/actions", handlerGetActions)
	http.HandleFunc("/get/install", handlerGetInstallInfo)
	http.HandleFunc("/get/images", handlerGetImageInfo)

	ticker := time.NewTicker(time.Minute * 2)
	tickerActions := time.NewTicker(time.Minute * 3)
	tickerTools := time.NewTicker(time.Minute * 3)
	tickerImages := time.NewTicker(time.Minute * 4)

	go func() {
		for t := range ticker.C {
			sw.AgregateInstalInfo()
			fmt.Println("Tick at", t)
		}
	}()
	go func() {
		for t := range tickerActions.C {
			sw.AgregateActions()
			fmt.Println("Tick actions at", t)
		}
	}()
	go func() {
		for t := range tickerTools.C {
			sw.AgregateTools()
			fmt.Println("Tick tools at", t)
		}
	}()
	go func() {
		for t := range tickerImages.C {
			sw.AgregateImageProps()
			fmt.Println("Tick image at", t)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
