package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	_ "github.com/lib/pq"
)

var data string
var db *sql.DB
var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/GoogleCallback",
		ClientID:     os.Getenv("googlekey"),
		ClientSecret: os.Getenv("googlesecret"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

//Request from krita
type Request struct {
	ApplicationVersion struct {
		Version string `json:"value"`
	} `json:"applicationVersion"`
	Compiler struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"compiler"`
	Locale struct {
		Language string `json:"language"`
	} `json:"locale"`
	Opengl struct {
		GlslVersion string `json:"glslVersion"`
		Renderer    string `json:"renderer"`
		Vendor      string `json:"vendor"`
	} `json:"opengl"`
	Platform struct {
		Os      string `json:"os"`
		Version string `json:"version"`
	} `json:"platform"`
	QtVersion struct {
		Version string `json:"value"`
	} `json:"qtVersion"`
	Screens []struct {
		Dpi    float64 `json:"dpi"`
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"screens"`
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

func insertToDb(req Request) {
	const insertQuery = "INSERT into generalinfo  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)"
	_, err := db.Exec(insertQuery, "userID1", req.ApplicationVersion.Version, req.Compiler.Version, req.Compiler.Type, req.Locale.Language, req.Opengl.GlslVersion, req.Opengl.Renderer, req.Opengl.Vendor, req.Platform.Os, req.Platform.Version, req.QtVersion.Version, int(req.Screens[0].Dpi), int(req.Screens[0].Height), int(req.Screens[0].Width))
	if err != nil {
		fmt.Println("insert error!")
	} else {
		fmt.Println("insert ended!")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("REQUEST")
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	data = data + "\n" + string(bodyBuffer)
	fmt.Printf(string(bodyBuffer))
	fmt.Printf("after parse")
	result := parse(bodyBuffer)
	fmt.Printf(string(result.QtVersion.Version))
	insertToDb(result)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Last requests</h1><div>%s</div>", data)
}
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Content: %s\n", contents)
}
func main() {
	fmt.Printf("hello")
	//	connectionString :=
	//	fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	var err error
	db, err = sql.Open("postgres", "user=root password=1111 dbname=root") //небезопасно, но пока сойдет
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/receiver/submit/org.krita.krita", handler)
	http.HandleFunc("/GoogleLogin", handleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	http.HandleFunc("/", viewHandler)

	http.ListenAndServe(":80", nil)
}
