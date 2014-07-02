package main

import (
	"github.com/googollee/go-socket.io"
	"log"
)

func test() {
	client, err := socketio.Dial("http://localhost:8012/")
	if err != nil {
		panic(err)
	}

	client.On("connect", func(ns *socketio.NameSpace, message string) {
		log.Println("connected")
	})

	client.On("news", func(ns *socketio.NameSpace, message string) {
		log.Println("news", message)
	})
	client.On("alerts", func(ns *socketio.NameSpace, message string) {
		log.Println("a", message)
	})

	client.Run()

}

func main() {
	go test()
}
