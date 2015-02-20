package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dat, _ := ioutil.ReadFile("index.html")
	tmpl := template.New("client")
	tmpl.Parse(string(dat))
	tmpl.Execute(w, "My firend")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8765", nil)
}
