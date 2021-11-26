package Handler

import (
	"net/http"
	"path"
	"text/template"
	"web/entitas"
)

func Putra(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(path.Join("view", "identitas.html"), path.Join("view", "layout.html"))

	if err != nil {
		panic(err)
	}

	person := entitas.Identitas{Nama: "putra", Umur: 22}

	if err = templ.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
