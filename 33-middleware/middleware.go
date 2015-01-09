package main

import "net/http"
import "fmt"
import "strings"

type SingleHost struct {
	handler     http.HandlerFunc
	allowedHost string
}

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")[0]

	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func NewSingleHost(handler http.HandlerFunc, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello you little motherfucker!"))
}

func main() {
	singleHosted := NewSingleHost(myHandler, "localhost")
	http.ListenAndServe(":8080", singleHosted)
}
