package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	fileserver := http.FileServer(http.Dir("assets"))

	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
