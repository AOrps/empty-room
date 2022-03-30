package frontend

import (
	"net/http"
	"text/template"
)

const (
	TMPLPARSE = "web/template/*.html"
)

func mapfx(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob(TMPLPARSE))
	navBar := NavBar()
	tmpl.ExecuteTemplate(w, "head", navBar)
	tmpl.ExecuteTemplate(w, "map", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)
}
