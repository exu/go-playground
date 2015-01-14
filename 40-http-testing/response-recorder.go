package main

import "fmt"
import "net/http/httptest"
import "net/http"
import "log"

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r)
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://l:7777", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
