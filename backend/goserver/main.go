package main

import (
	"fmt"
	"io/ioutil"
	serv "kritaServers/backend/goserver/server"
	agr "kritaServers/backend/goserver/server/agregate"
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
	serv.InsertGeneralInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func handlerTools(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))
	serv.InsertToolInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func handlerImageProperties(w http.ResponseWriter, r *http.Request) {
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))
	serv.InsertImageInfo(bodyBuffer)
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
	serv.InsertActionInfo(bodyBuffer)
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
func handlerGetTools(w http.ResponseWriter, r *http.Request) {
	temp := agr.Agregated("tools")
	fmt.Fprintf(w, temp)
}
func handlerGetActions(w http.ResponseWriter, r *http.Request) {
	temp := agr.Agregated("actions")
	fmt.Fprintf(w, temp)
}
func handlerGetInstallInfo(w http.ResponseWriter, r *http.Request) {
	type1 := r.URL.Query().Get("type")
	if len(type1) != 0 {
		dataOfType := agr.AgregatedInstall(type1)
		fmt.Fprintf(w, dataOfType)
		fmt.Println(dataOfType)
		return
	}
	temp := agr.Agregated("install")
	fmt.Println(temp)
	fmt.Fprintf(w, temp)
}
func handlerGetImageInfo(w http.ResponseWriter, r *http.Request) {
	type1 := r.URL.Query().Get("type")
	if len(type1) != 0 {
		dataOfType := agr.AgregatedInstall(type1)
		fmt.Fprintf(w, dataOfType)
		return
	}
	temp := agr.Agregated("images")
	fmt.Fprintf(w, temp)
}

func main() {
	fmt.Printf("hello")
	serv.InitDB()
	defer serv.Session.Close()

	http.HandleFunc("/install/receiver/submit/org.krita.krita", handlerInstall)
	http.HandleFunc("/tools/receiver/submit/org.krita.krita", handlerTools)

	http.HandleFunc("/imageProperties/receiver/submit/org.krita.krita", handlerImageProperties)
	http.HandleFunc("/asserts/receiver/submit/org.krita.krita", handlerAsserts)
	http.HandleFunc("/actions/receiver/submit/org.krita.krita", handlerActions)

	http.HandleFunc("/GoogleLogin", serv.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", serv.HandleGoogleCallback)

	http.HandleFunc("/get/tools", handlerGetTools)
	http.HandleFunc("/get/actions", handlerGetActions)
	http.HandleFunc("/get/install", handlerGetInstallInfo)
	http.HandleFunc("/get/images", handlerGetImageInfo)

	ticker := time.NewTicker(time.Minute * 2)
	tickerActions := time.NewTicker(time.Minute * 3)
	tickerTools := time.NewTicker(time.Minute * 3)
	tickerImages := time.NewTicker(time.Minute * 4)

	go func() {
		for _ = range ticker.C {
			agr.AgregateInstalInfo()
			//	fmt.Println("Tick at", t)
		}
	}()
	go func() {
		for _ = range tickerActions.C {
			agr.AgregateActions()
			//fmt.Println("Tick actions at", t)
		}
	}()
	go func() {
		for _ = range tickerTools.C {
			agr.AgregateTools()
			//	fmt.Println("Tick tools at", t)
		}
	}()
	go func() {
		for _ = range tickerImages.C {
			agr.AgregateImageProps()
			//fmt.Println("Tick image at", t)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
