package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Name   string
	Author string
	Genre  string
	Year   int
	Rating int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "books/books.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT Name, Author, Genre, Year, Rating FROM Books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	books := []Book{}

	_, err = db.Exec("insert into books (Name, Author, Genre, Year, Rating) values ('War and Peace', 'Leo Tolstoy', 'Historical novel', 1865, 7)")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Name, &book.Author, &book.Genre, &book.Year, &book.Rating)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, book)
	}
	fmt.Println(books)

}
