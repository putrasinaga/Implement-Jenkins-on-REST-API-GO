package Handler

import (
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func TestCss(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles(path.Join("view", "css.html"), path.Join("view", "layout.html"))
	if err != nil {
		panic(err)
	}

	templ.Execute(w, nil)
}

func Nomor(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	numb, err := strconv.Atoi(id)

	if err != nil || numb < 1 {

		http.NotFound(w, r)
		return

	}

	// fmt.Fprintln(w, "ini adalah nomor : ", numb)

	templ, err := template.ParseFiles(path.Join("view", "product.html"), path.Join("view", "layout.html"))
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"title": "halaman produk ",
		"nomor": numb,
	}

	err = templ.Execute(w, data)
	if err != nil {
		panic(err)
	}

}

func Route(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("ini adalah home"))

	var data = map[string]interface{}{
		"title":   "ini adalah hal baru",
		"content": "contoh ajaa",
	}

	var tmpl, err = template.ParseFiles(path.Join("view", "index.html"), path.Join("view", "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
