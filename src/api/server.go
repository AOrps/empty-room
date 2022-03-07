package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

const (
	PORT         = 9993
	JSONCONFPATH = "out.json"
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
	content, err := ioutil.ReadFile(filename)
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

	// Map values in filename (json file) to `dat` var
	fpath, err := filepath.Abs(JSONCONFPATH)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	JSONGo(fpath, &dat)

	bset := building_set(dat)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			result := make(map[string][]string)

			if err := r.ParseForm(); err != nil {
				log.Fatal(err.Error())
			}
			building := r.FormValue("building")
			day := r.FormValue("day")
			log.Printf("<POST: building:[%s]  day:[%s]>", building, day)

			if building == "CC" {
				building = "CTR"
			}

			if bset[building] {
				rooms := listRooms(dat, building)

				for _, room := range rooms {
					classesNum := len(dat[building][room][day])

					for i := 0; i < classesNum; i++ {
						result[room] = append(result[room], dat[building][room][day][i].Time)
					}
				}
			}
			// outputs result from out.json in json when ppl connect to it
			enc := json.NewEncoder(w)
			enc.Encode(result)

		default:
			fmt.Fprintf(w, "POST plz")
		}
	})

	fmt.Printf("Server: 'http://localhost:%s'\n", port)
	http.ListenAndServe(":"+port, nil)
}
