package main

import (
	"fenye/conrtrol"
	"net/http"
)

func main() {
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc(`/`, conrtrol.Viewdemo)
	http.HandleFunc(`/page`, conrtrol.ArtiPage)

	http.ListenAndServe(":6060", nil)
}
