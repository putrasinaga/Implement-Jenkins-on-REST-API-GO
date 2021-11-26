package Handler

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

func Form(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		templ, err := template.ParseFiles(path.Join("view", "form.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "terjadi kesalahan", 500)
			return // jika ada yang error maka code akan berjalan sampai sini
		}

		templ.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "terjadi kesalahan", 500)
			return
		}

		return

	}

	http.Error(w, "method hanya menerima GET", 404)

}
