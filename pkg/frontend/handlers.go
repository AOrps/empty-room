package frontend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	pjson "github.com/AOrps/empty-room/pkg/parsejson"
)

const (
	TMPLPARSE = "web/template/*.html"
)

// Map:
func Map(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "map", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

// Schedule:
func Schedule(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "sched", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

// About:
func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "about", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}

// FindRoom:
func FindRoom(w http.ResponseWriter, r *http.Request) {

	var dat pjson.Info

	// fpath, err := filepath.Abs("web/api/out.json")
	fpath := "web/api/out.json"
	// if err != nil {
	// 	log.Fatalf("%s", err.Error())
	// }

	pjson.JSONGo(fpath, &dat)

	bset := pjson.BuildingSet(dat)

	switch r.Method {
	case "POST":
		result := make(map[string][]string)

		if err := r.ParseForm(); err != nil {
			log.Fatal(err.Error())
		}

		building := strings.ToUpper(r.FormValue("building"))
		day := strings.ToUpper(r.FormValue("day"))

		// data := url.Values{
		// 	"building": {r.FormValue("building")},
		// 	"day":      {r.FormValue("day")},
		// }

		fmt.Fprintf(w, "%s -> %s ", day, building)

		if building == "CC" {
			building = "CTR"
		}

		weekDay := time.Now().Weekday()

		switch strings.ToLower(day) {
		case "monday":
			day = "M"
		case "tuesday":
			day = "T"
		case "wednesday":
			day = "W"
		case "thursday":
			day = "R"
		case "friday":
			day = "F"
		case "saturday":
			day = "S"
		default:
			dayz := []string{"SUN", "M", "T", "W", "R", "F", "S"}
			day = dayz[int(weekDay)]
		}

		if bset[building] {
			rooms := pjson.ListRooms(dat, building)

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
		// resp, err := http.PostForm(os.Getenv("BACKEND"), data)
		// u.Check(err)
		// defer resp.Body.Close()
		// body, err := ioutil.ReadAll(resp.Body)
		// u.Check(err)

		// fmt.Printf("%v", string(body))
		// fmt.Fprintf(w, "%v", string(body))

		// body is []byte
		// requ.POSTRESPONSE(w, body)

	default:
		tmpl := template.Must(template.ParseGlob(TMPLPARSE))
		navBar := NavBar()
		tmpl.ExecuteTemplate(w, "head", navBar)

		tmpl.ExecuteTemplate(w, "find-start", nil)

		// Info Block (Table with Info about Buildings)
		infoblock := JsonInfo()
		tmpl.ExecuteTemplate(w, "info", infoblock)

		// Form
		tmpl.ExecuteTemplate(w, "find-form", nil)

		tmpl.ExecuteTemplate(w, "find-end", nil)
		tmpl.ExecuteTemplate(w, "footer", nil)
	}
}
