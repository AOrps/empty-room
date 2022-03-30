package parsejson

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
func ListRooms(data Info, building string) []string {
	var retSlice []string
	for room := range data[building] {
		retSlice = append(retSlice, room)
	}
	return retSlice
}

// buildingSet: creates a set with only buildings
func BuildingSet(data Info) map[string]bool {
	set := make(map[string]bool)
	for building := range data {
		set[building] = true
	}
	return set
}
