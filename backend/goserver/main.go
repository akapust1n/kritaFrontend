package main

import (
	"fmt"
	"io/ioutil"
	sw "kritaServers/backend/goserver/server"
	"net/http"
	"time"

	_ "github.com/lib/pq"
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
	sw.InsertGeneralInfo(bodyBuffer)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Last requests</h1><div>%s</div>", "in console")
}

func main() {
	fmt.Printf("hello")
	//	connectionString :=
	//	fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	sw.InitDB()
	defer sw.Db.Close()

	http.HandleFunc("/receiver/submit/org.krita.krita", handler)
	http.HandleFunc("/GoogleLogin", sw.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", sw.HandleGoogleCallback)
	http.HandleFunc("/agregatedData", sw.AgregatedDataHandler)

	http.HandleFunc("/", viewHandler)

	//ticker := time.NewTicker(time.Minute * 2)
	ticker := time.NewTicker(time.Minute * 2)

	go func() {
		for t := range ticker.C {
			sw.CollectData()
			fmt.Println("Tick at", t)
		}
	}()
	http.ListenAndServe(":8080", nil)
}
