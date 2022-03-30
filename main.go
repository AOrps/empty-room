package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	fe "github.com/AOrps/empty-room/pkg/frontend"
	"github.com/joho/godotenv"
)

const (
	PORT      = 9994
	TMPLPARSE = "web/template/*.html"
)

// init: function that get's called on initialization
func init() {
	// loads the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := strconv.Itoa(PORT)

	fmt.Printf("http://localhost:%s\n", port)

	fs := http.FileServer(http.Dir("./web/"))
	// Puts everything from File Server into a /static/ directory
	// puts prefix of /static/
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", fe.FindRoom)
	http.HandleFunc("/map", fe.Map)
	http.HandleFunc("/schedule", fe.Schedule)
	http.HandleFunc("/find-room", fe.FindRoom)
	http.HandleFunc("/about", fe.About)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
