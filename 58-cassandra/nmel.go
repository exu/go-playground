/* Before you execute the program, Launch `cqlsh` and execute:

CREATE KEYSPACE test1 WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;
USE test1;
CREATE TYPE answer_task( euid bigint, result int);
CREATE TABLE attempt( id timeuuid, user_assignment_id int, answers list <frozen<answer_task>>, primary KEY (user_assignment_id, id) );
INSERT INTO attempt (user_assignment_id, id, answers) VALUES (1, now(), [{euid:12,result:1}]);
INSERT INTO attempt (user_assignment_id, id, answers) VALUES (2, now(), [{euid:1332,result:2}]);

*/
package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type Answer struct {
	Euid      int `cql:"euid"`
	Questions int `cql:"questions"`
	Error     int `cql:"error"`
	Result    int `cql:"result"`
	MaxResult int `cql:"max_result"`
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
	cluster.ProtoVersion = 3
	cluster.Consistency = gocql.Quorum
	session, _ = cluster.CreateSession()

	return session
}

func main() {
	session := newSession()
	defer session.Close()

	attempt := Attempt{}
	if err := session.Query(`SELECT id, user_assignment_id, answers FROM attempt`).
		Consistency(gocql.One).
		Scan(&attempt.Id, &attempt.UserAssignmentId, &attempt.Answers); err != nil {
		log.Fatal(err)
	}

	fmt.Println("attempt:", attempt)
}
