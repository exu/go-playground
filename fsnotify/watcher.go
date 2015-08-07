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

	err = watcher.Watch("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	// in original example it was wrapped in go func(){}
	// I prefer blocking version
	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
