package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 9993
)

/*
POST API:
	- Building
	- Day
ex:
POST
- building:GITC
- day:M
*/

// Structures
type Detail struct {
	Key        string `json:"key"`
	Title      string `json:"title"`
	Time       string `json:"time"`
	Instructor string `json:"instructor"`
}

// Info -> Building: Room: Day: Details
type Info map[string]map[string]map[string][]Detail

// JSONGo: reads and unmarhsals a json file to an object reference
func JSONGo(filename string, data *Info) {
	content, err := ioutil.ReadFile("out.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, data)
	if err != nil {
		log.Fatal(err)
	}
}

// listRooms: list rooms that are used as classrooms from a building
func listRooms(data Info, building string) []string {
	var retSlice []string
	for room := range data[building] {
		retSlice = append(retSlice, room)
	}
	return retSlice
}

// building_set: creates a set with only buildings
func building_set(data Info) map[string]bool {
	set := make(map[string]bool)
	for building := range data {
		set[building] = true
	}
	return set
}

func main() {
	port := strconv.Itoa(PORT)
	var dat Info

	fs := http.FileServer(http.Dir("."))

	http.Handle("/data/", http.StripPrefix("/data", fs))

	// Map values in filename (json file) to `dat` var
	JSONGo("data/out.json", &dat)

	bset := building_set(dat)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				log.Fatal(err.Error())
			}
			building := r.FormValue("building")
			day := r.FormValue("day")
			log.Printf("{-->  POST: building:[%s]  day:[%s]}", building, day)

			if building == "CC" {
				building = "CTR"
			}

			if bset[building] {
				rooms := listRooms(dat, building)
				fmt.Fprintf(w, "{%s}\n", day)

				for _, room := range rooms {
					// fmt.Println(room)
					classesNum := len(dat[building][room][day])
					for i := 0; i < classesNum; i++ {
						fmt.Fprintf(w, "%s -> %v\n", room, dat[building][room][day][i].Time)
					}
				}
			}
			// log.Println(building)

		default:
			fmt.Fprintf(w, "POST plz")
		}
	})

	fmt.Printf("Server: 'http://localhost:%s'\n", port)
	http.ListenAndServe(":"+port, nil)
}
