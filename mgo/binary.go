package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sample struct {
	Name string `bson:"name"`
	Bin  []byte `bson:"bin"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("bin")
	c.DropCollection()

	err = c.Insert(&Sample{Name: "CoreSample1", Bin: []byte("Some random goop")})
	if err != nil {
		panic(err)
	}

	result := Sample{}
	err = c.Find(bson.M{"name": "CoreSample1"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
