package main

import (
	"net/http"

	"encoding/json"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"log"
)

func Json(w http.ResponseWriter, item interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(item)

	if err != nil {
		log.Println(err)
		return
	}

	w.Write(response)
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	Json(w, &map[string]interface{}{"aaa": 1})
}

func main() {
	goji.Get("/ping", hello)
	goji.Serve()
}
