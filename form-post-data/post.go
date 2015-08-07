package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// when x-www-form-urlencoded
	postFormValue := r.PostFormValue("a")
	// otherwise
	formValue := r.FormValue("a")

	fmt.Fprintf(
		w,
		"Post form value: %v , form value: %v\n",
		postFormValue,
		formValue,
	)

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7890", nil)
}
