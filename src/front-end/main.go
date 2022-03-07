package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	L "github.com/AOrps/empty-room/src/front-end/lib"
)

const (
	PORT = 9994
	API  = "http://ec2-18-119-118-48.us-east-2.compute.amazonaws.com/api"
)

func mapfx(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	navBar := L.NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "map", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
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

		resp, err := http.PostForm(API, data)
		L.Check(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		L.Check(err)

		// fmt.Printf("%v", string(body))
		fmt.Fprintf(w, "%v", string(body))
	default:
		tmpl := template.Must(template.ParseGlob("templates/*.html"))
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
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
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
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", findRoom)
	http.HandleFunc("/map", mapfx)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/find-room", findRoom)
	http.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
