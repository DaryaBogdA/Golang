package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type book struct {
	id_book      int
	name_book    string
	author       string
	year_realize int
	id           int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(MySQL-8.2:3306)/library")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := []book{}
	for rows.Next() {
		b := book{}
		err := rows.Scan(&b.id, &b.name_book, &b.author, &b.year_realize, &b.id_book)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, b)
	}

	fmt.Println(books)
}
