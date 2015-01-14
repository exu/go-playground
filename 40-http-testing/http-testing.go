package main

import "fmt"
import "net/http/httptest"
import "log"
import "net/http"
import "io/ioutil"

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client !!")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(greeting) == "Hello, client !!\n" {
		fmt.Printf("%s", greeting)
	} else {
		log.Fatal("Not greeting at all!! ")
	}

}
