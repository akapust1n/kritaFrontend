package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var data string

//Request from krita
type Request struct {
	Foow struct {
		Bar string `json:"value"`
	} `json:"qtVersion"`
	//  FooBar  string `json:"foo.bar"`
}

func parse(req []byte) Request {
	var result Request
	err := json.Unmarshal([]byte(req), &result)
	if err == nil {
		fmt.Printf("%+v\n", result)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", result)
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "KRITA REQUEST ")
	fmt.Printf("REQUEST")
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("ss")
	data = data + "\n" + string(bodyBuffer)
	fmt.Printf(string(bodyBuffer))
	fmt.Printf("after parse")
	parse(bodyBuffer)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Last requests</h1><div>%s</div>", data)
}

func main() {
	http.HandleFunc("/receiver/submit/org.krita.krita", handler)
	http.HandleFunc("/", viewHandler)
	//var pro Request
	// s := `{"qtVersion":{ "value": "5.2.0"}}`
	// parse([]byte(s), &pro)

	http.ListenAndServe(":8080", nil)
}
