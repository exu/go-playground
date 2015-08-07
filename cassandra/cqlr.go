package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/relops/cqlr"
)

type Answer struct {
	Euid int `cql:"euid"`
}

type Attempt struct {
	Id               gocql.UUID `cql:"id"`
	UserAssignmentId string     `cql:"user_assignment_id"`
	Answers          []Answer   `cql:"answers"`
}

func newSession() *gocql.Session {
	var session *gocql.Session

	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "nmel"
	cluster.Consistency = gocql.Quorum
	session, _ = cluster.CreateSession()

	return session
}

func main() {

	session := newSession()
	defer session.Close()

	// iter := session.Query(
	// 	`SELECT * FROM attempt WHERE id = ?`,
	// 	"f7b14670-1af8-11e5-9608-6d8f944ad4f3",
	// ).Iter()

	// data, _ := iter.RowData()
	// fmt.Println("Attempt:", data)

	query := session.Query(`SELECT id, user_assignment_id, answers FROM attempt`)
	bind := cqlr.BindQuery(query)

	attempt := Attempt{}
	for bind.Scan(&attempt) {
		fmt.Println(attempt)
	}

	// if err := iter.Close(); err != nil {
	// 	fmt.Println(err)
	// }

}
