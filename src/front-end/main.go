package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	
	L "github.com/AOrps/empty-room/src/front-end/lib"
)

const (
	PORT = 9994
	API  = "https://ec2-18-119-118-48.us-east-2.compute.amazonaws.com"
)

func mapfx(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	navBar := []string{"map", "schedule", "find-room"}
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "map", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

func schedule(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	navBar := []string{"map", "schedule", "find-room"}
	tmpl.ExecuteTemplate(w, "head", navBar)

	// Exec

	tmpl.ExecuteTemplate(w, "footer", nil)
}

func findRoom(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	navBar := []string{"map", "schedule", "find-room"}
	tmpl.ExecuteTemplate(w, "head", navBar)

	tmpl.ExecuteTemplate(w, "find-start", nil)


	// Exec
	infoblock := L.JsonInfo()
	tmpl.ExecuteTemplate(w, "info", infoblock)
	tmpl.ExecuteTemplate(w, "find-form", nil)


	tmpl.ExecuteTemplate(w, "find-end", nil)


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

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
