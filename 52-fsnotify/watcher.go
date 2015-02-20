package main

import (
	"golang.org/x/exp/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}

	// in example it was wrapped in go func(){}
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
