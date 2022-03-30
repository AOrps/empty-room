package parsejson

// Structures
type Detail struct {
	Key        string `json:"key"`
	Title      string `json:"title"`
	Time       string `json:"time"`
	Instructor string `json:"instructor"`
}

// Info -> Building: Room: Day: Details
type Info map[string]map[string]map[string][]Detail
