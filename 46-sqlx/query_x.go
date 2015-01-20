package main

import "fmt"
import "github.com/jmoiron/sqlx"

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Item struct {
	Id          int            `db:"id"`
	Name        sql.NullString `db:"nazwa"`
	Description sql.NullString `db:"opis"`
}

func main() {
	db := sqlx.MustConnect("mysql", "root:p-o0i9@tcp(l:3306)/test")
	fmt.Println(db)

	rows, err := db.Queryx("SELECT id, nazwa, opis FROM przepisy")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var item Item
		err = rows.StructScan(&item)

		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"%d - %s:  %s\n\n===================\n\n",
			item.Id,
			item.Name.String,
			item.Description.String,
		)
	}
}
