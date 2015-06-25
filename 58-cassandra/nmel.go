/* Before you execute the program, Launch `cqlsh` and execute:

CREATE KEYSPACE nmel WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;
CREATE TYPE answer_value( value varchar, error boolean );
CREATE TYPE answer_question ( euid bigint , values list <frozen<answer_value>>, error boolean, result int, max_result int );
CREATE TYPE answer_task( euid bigint, questions list <frozen<answer_question>>, error boolean, result int, max_result int );
CREATE TABLE attempt( id timeuuid, user_assignment_id int, answers list <frozen<answer_task>>, attempt_date timestamp, time_spent int, saved boolean, submitted boolean, graded boolean, comments list <varchar>, primary KEY (user_assignment_id, id) );

*/
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type Answer struct {
	Euid      int `cql:"euid"`
	Questions int `cql:"questions"`
	Error     int `cql:"error"`
	Result    int `cql:"result"`
	MaxResult int `cql:"max_result"`
}

type Attempt struct {
	Id               gocql.UUID    `cql:"id"`
	UserAssignmentId string        `cql:"user_assignment_id"`
	Answers          []interface{} `cql:"answers"`
}

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "nmel"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert row
	// timeUUID := gocql.TimeUUID()
	// answer := Answer{Euid: 122}
	// if err := session.Query(`INSERT INTO attempt (user_assignment_id, id, answers) VALUES (?, ?, ?)`,
	// 	3, timeUUID, answer).Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	var id gocql.UUID
	var uaid int
	var answer Answer

	if err := session.Query(`SELECT id, user_assignment_id, answers FROM attempt`).
		Consistency(gocql.One).
		Scan(&id, &uaid, &answer); err != nil {
		log.Fatal(err)
	}
	fmt.Println("a:", id, uaid)

	// // list all tweets
	// iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, id).Iter()
	// for iter.Scan(&id, &text) {
	// 	fmt.Println("Tweet:", id, text)
	// }
	// if err := iter.Close(); err != nil {
	// 	log.Fatal(err)
	// }
}
