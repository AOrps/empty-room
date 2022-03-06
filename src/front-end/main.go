package main

import (
	"net/http"
	"strconv"
)

const (
	PORT = 9994
	API  = "https://ec2-18-119-118-48.us-east-2.compute.amazonaws.com"
)

func root(w http.ResponseWriter, r *http.Request) {

}

func map() {}

func schedule() {}

func findRoom() {}




func main() {
	port := strconv.Itoa(PORT)

	http.HandleFunc("/", root)

	http.ListenAndServe(":"+port, nil)
}
