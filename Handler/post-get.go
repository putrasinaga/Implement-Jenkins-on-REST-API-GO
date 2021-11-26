package Handler

import (
	"net/http"
)

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah Get"))
	case "POST":
		w.Write([]byte("ini adalah Post"))
	default:
		http.Error(w, "ada terjadi sesuat", http.StatusBadRequest)
	}
}
