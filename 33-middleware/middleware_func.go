package main

import "net/http"
import _ "fmt"
import "strings"

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello you little motherfucker!"))
}

func singleHostMiddleware(handler http.HandlerFunc, allowedHost string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := strings.Split(r.Host, ":")[0]
		if host == allowedHost {
			w.Header().Add("X-Host-Allowed", "true")
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Add("X-Host-Allowed", "false")
			w.WriteHeader(403)
		}
	})
}

func main() {
	http.ListenAndServe(":8080", singleHostMiddleware(myHandler, "localhost"))
}
