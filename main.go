package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	L "github.com/AOrps/empty-room/src/front-end/lib"
	"github.com/joho/godotenv"
)

const (
	PORT      = 9994
	TMPLPARSE = "web/template/*.html"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func mapfx(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := L.NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "map", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := L.NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)

	// Exec
	tmpl.ExecuteTemplate(w, "sched", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)
}

func findRoom(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		data := url.Values{
			"building": {r.FormValue("building")},
			"day":      {r.FormValue("day")},
		}

		resp, err := http.PostForm(os.Getenv("BACKEND"), data)
		L.Check(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		L.Check(err)

		// fmt.Printf("%v", string(body))
		// fmt.Fprintf(w, "%v", string(body))

		// body is []byte
		L.POSTRESPONSE(w, body)

	default:
		tmpl := template.Must(template.ParseGlob(TMPLPARSE))
		navBar := L.NavBar()
		tmpl.ExecuteTemplate(w, "head", navBar)

		tmpl.ExecuteTemplate(w, "find-start", nil)

		// Info Block (Table with Info about Buildings)
		infoblock := L.JsonInfo()
		tmpl.ExecuteTemplate(w, "info", infoblock)

		// Form
		tmpl.ExecuteTemplate(w, "find-form", nil)

		tmpl.ExecuteTemplate(w, "find-end", nil)
		tmpl.ExecuteTemplate(w, "footer", nil)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := L.NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)

	// Exec
	tmpl.ExecuteTemplate(w, "about", nil)

	tmpl.ExecuteTemplate(w, "footer", nil)
}

func main() {
	port := strconv.Itoa(PORT)

	fmt.Printf("http://localhost:%s\n", port)

	fs := http.FileServer(http.Dir("."))
	// Puts everything from File Server into a /assets/ directory
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", findRoom)
	http.HandleFunc("/map", mapfx)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/find-room", findRoom)
	http.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
