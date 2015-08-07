package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Users  []string
	Groups []string
}

func main() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users)  // output: [UserA, UserB]
	fmt.Println(configuration.Groups) // output: [GroupA]
}
