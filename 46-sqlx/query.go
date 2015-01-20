package main

import "fmt"
import "github.com/jmoiron/sqlx"

// import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db := sqlx.MustConnect("mysql", "root:p-o0i9@tcp(l:3306)/test")
	fmt.Println(db)

	rows, err := db.Query("SELECT id, nazwa, opis FROM przepisy")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var nazwa string
		var opis string

		err = rows.Scan(&id, &nazwa, &opis)

		if err != nil {
			panic(err)
		}

		fmt.Println(id, nazwa, opis)
	}
}
