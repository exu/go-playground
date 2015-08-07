package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Starting file serving\n")

	fileValue, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(
		w,
		"File: %v\nHeader: %v\n",
		*fileValue,
		fileHeader,
	)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7890", nil)
}
