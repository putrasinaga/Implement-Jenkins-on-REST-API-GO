package Handler

import (
	"net/http"
	"path"
	"text/template"
	"web/entitas"
)

func User(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(path.Join("view", "user.html"), path.Join("view", "layout.html"))

	if err != nil {
		panic(err)
	}

	user := []entitas.User{
		{Id: 1, Nama: "putra", Umur: 10},
		{Id: 2, Nama: "saut", Umur: 22},
		{Id: 3, Nama: "sinaga", Umur: 16},
	}

	if err = templ.Execute(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
