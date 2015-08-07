package main

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/hailocab/gocassa"
)

type Answer struct {
	Euid      int   `cql:"euid"`
	Questions []int `cql:"questions"`
}

type Attempt struct {
	Id               gocql.UUID `cql:"id"`
	UserAssignmentId string     `cql:"user_assignment_id"`
	Answers          []Answer   `cql:"anwsers"`
}

func main() {
	keySpace, err := gocassa.ConnectToKeySpace("nmel", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}

	attemptTable := keySpace.MapTable("attempt", "id", Attempt{})
	result := Attempt{}
	attemptTable.Read("f7b14670-1af8-11e5-9608-6d8f944ad4f3", &result).Run()
	fmt.Println(result)

	// attemptTable := keySpace.Table("attempt", Attempt{}, gocassa.Keys{
	// 	PartitionKeys: []string{"Id"},
	// })
	// err = attemptTable.Set(Sale{
	// 	Id:         "sale-1",
	// 	CustomerId: "customer-1",
	// 	SellerId:   "seller-1",
	// 	Price:      42,
	// 	Created:    time.Now(),
	// }).Run()

	// if err != nil {
	// 	panic(err)
	// }

	// result := Attempt{}
	// if err := attemptTable.Where(gocassa.Eq("id", "sale-1")).ReadOne(&result).Run(); err != nil {
	// 	panic(err)
	// }
	fmt.Println(result)
}
