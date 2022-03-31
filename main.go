package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	fe "github.com/AOrps/empty-room/pkg/frontend"
	pjson "github.com/AOrps/empty-room/pkg/parsejson"
	"github.com/joho/godotenv"
)

const (
	PORT      = 9994
	TMPLPARSE = "web/template/*.html"
	JSONAPI   = "web/api/out.json"
)

// init: function that get's called on initialization
func init() {
	// loads the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Api(w http.ResponseWriter, r *http.Request) {
	var data pjson.Info
	pjson.JSONGo(JSONAPI, &data)
	// fmt.Fprintf(w, "%s", ":"+strconv.Itoa(PORT+1))
	enc := json.NewEncoder(w)
	// A nice pretty print of JSON via SetIdent to values
	enc.SetIndent("", "  ")
	enc.Encode(data)
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
	http.HandleFunc("/api", fe.Api)

	// Api Server
	apiServer := http.NewServeMux()
	apiServer.HandleFunc("/", Api)
	go func() {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT+1), apiServer))
	}()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
