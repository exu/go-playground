package main

import "fmt"
import "github.com/jmoiron/sqlx"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Item struct {
	Id            int            `db:"id"`
	Name          sql.NullString `db:"nazwa"`
	Description   sql.NullString `db:"opis"`
	Sklad         sql.NullString `db:"sklad"`
	Przygotowanie sql.NullString `db:"przygotowanie"`
	Tags          sql.NullString `db:"tags"`
	CategoryId    sql.NullInt64  `db:"category_id"`
	Category2Id   sql.NullInt64  `db:"category2_id"`
	Category3Id   sql.NullInt64  `db:"category3_id"`
	Created       sql.NullString `db:"created"`
	Updated       sql.NullString `db:"updated"`
	PublishAt     sql.NullString `db:"publish_at"`
	CookTime      sql.NullString `db:"czas_przygotowania"`
	IsMain        bool           `db:"main"`
	Type          sql.NullString `db:"type"`
	Video         sql.NullString `db:"video"`
}

func main() {
	db := sqlx.MustConnect("mysql", "root:p-o0i9@tcp(l:3306)/test")
	var item Item
	err := db.Get(&item, "SELECT * FROM przepisy ORDER BY ID DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}
