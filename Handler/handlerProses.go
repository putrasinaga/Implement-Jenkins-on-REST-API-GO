package Handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Proses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "ada yang salah", 404)
			return
		}

		Firstname := r.Form.Get("firstname")
		Lastname := r.Form.Get("lastname")

		templ, err := template.ParseFiles(path.Join("view", "result.html"))

		if err != nil {
			log.Print(err)
			http.Error(w, "ada yang salah pada konversi html", 404)
			return

		}

		data := map[string]interface{}{

			"firstname": Firstname,
			"lastname":  Lastname,
		}

		templ.Execute(w, data)
		if err != nil {
			log.Print(err)
			http.Error(w, "eksekusi gagal", 404)
			return
		}

		return
	}

	http.Error(w, "tidak ada opsi itu", 404)
}
