package main

import (
	"fmt"
	"net/http"
	"web/Handler"
)

func main() {

	http.HandleFunc("/index", Handler.TestCss)
	http.HandleFunc("/putra", Handler.Putra)
	http.HandleFunc("/nomor", Handler.Nomor)
	http.HandleFunc("/user", Handler.User)
	http.HandleFunc("/", Handler.Route)
	http.HandleFunc("/post-get", Handler.PostGet)
	http.HandleFunc("/form", Handler.Form)
	http.HandleFunc("/proses", Handler.Proses)

	fileserver := http.FileServer(http.Dir("assets"))

	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
