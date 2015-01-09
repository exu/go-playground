package main

import "os"
import "log"
import "io/ioutil"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	log.Println(err, string(bytes))
}
