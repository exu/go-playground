package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  e: 3
  f:
    - lalal
    - lolol
    - gogogo
      - lalala
`

type T struct {
	A string
	B struct {
		C int
		D []int ",flow"
	}
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	data, err := ioutil.ReadFile("parameters.yml")
	if err != nil {
		log.Fatal(err)
	}

	params := make(map[interface{}]map[string]string)
	err = yaml.Unmarshal([]byte(data), &params)
	fmt.Printf("--- params dump:\n%#v\n\n", (params))

}
