package lib

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type InfoBlock struct {
	Building string `json:"building"`
	Num      string `json:"num"`
	Abbrev   string `json:"abbrev"`
}

type Room struct {
	Room  string
	Times []string
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func JsonInfo() []InfoBlock {
	var data []InfoBlock
	jsonString := `[{"building":"Campus Center","num":"12","abbrev":"CTR | CC"},{"building":"Central Avenue Building","num":"6","abbrev":"CAB"},{"building":"Central King Building","num":"7","abbrev":"CKB"},{"building":"Colton Hall","num":"5","abbrev":"COLT"},{"building":"Cullimore Hall","num":"10","abbrev":"CULM"},{"building":"Electrical and Computer Eng. Center","num":"16","abbrev":"ECEC"},{"building":"Faculty Memorial Hall","num":"15","abbrev":"FMH"},{"building":"Fenster Hall","num":"8","abbrev":"FENS"},{"building":"Guttenberg Information Technology Center","num":"31","abbrev":"GITC"},{"building":"Kupfrian Hall","num":"13","abbrev":"KUPF"},{"building":"Mechanical Engineering Center","num":"30","abbrev":"ME"},{"building":"Tiernan Hall","num":"26","abbrev":"TIER"},{"building":"Weston Hall","num":"4","abbrev":"WEST"}]`

	json.Unmarshal([]byte(jsonString), &data)

	return data
}

func NavBar() []string {
	return []string{"map", "schedule", "find-room", "about"}
}

func POSTRESPONSE(w http.ResponseWriter, rawjson []byte) {

	var rooms map[string][]string

	json.Unmarshal([]byte(rawjson), &rooms)

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	navBar := NavBar()

	tmpl.ExecuteTemplate(w, "head", navBar)

	// Exec
	// fmt.Fprintf(w, "%v", rooms)
	tmpl.ExecuteTemplate(w, "rooms", rooms)

	tmpl.ExecuteTemplate(w, "footer", nil)
}
