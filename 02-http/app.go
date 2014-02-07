package main

import "net/http"
import "fmt"

func getStatus(status chan string, url string) {
	res, _ := http.Get(url)
	status <- res.Status + " " + url
}

func printStatus(status chan string) {
	for {
		fmt.Println(<-status)
	}
}

func main() {
	var status chan string = make(chan string, 3)

	uris := []string{
		"http://kasia.in/rss",
		"http://google.pl",
		"http://qarson.fr",
		"http://edpauto.fr",
		"http://kasia.in/rss",
		"http://google.pl",
		"http://qarson.fr",
		"http://edpauto.fr",
		"http://kasia.in/rss",
		"http://google.pl",
		"http://qarson.fr",
		"http://edpauto.fr",
		"http://kasia.in/rss",
		"http://google.pl",
		"http://qarson.fr",
		"http://edpauto.fr",
		"http://kasia.in/rss",
		"http://google.pl",
		"http://qarson.fr",
		"http://edpauto.fr",
	}

	for _, uri := range uris {
		go getStatus(status, uri)
	}

	go printStatus(status)

	var i string
	fmt.Scanln(&i)
}
