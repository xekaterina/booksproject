package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	entries, err := os.ReadDir("./books")
	if err != nil {
		log.Fatal(err)
	}

	books := make([]Book, 0, len(entries))

	for _, e := range entries {

		name := e.Name()

		dat, err := os.ReadFile("books/" + name)
		check(err)

		s := strings.Split(string(dat), "\n")
		fmt.Println(s[1])

		year, err := strconv.Atoi(s[3])
		if err != nil {
			// ... handle error
			panic(err)
		}

		rating, err := strconv.Atoi(s[4])
		if err != nil {
			// ... handle error
			panic(err)
		}

		var book = Book{Name: name, Author: s[1], Genre: s[2], Year: year, Rating: rating}
		books = append(books, book)
	}
	fmt.Println(books)

	// books := make([]Book, 0, 5)
	// dat, err := os.ReadFile("books/test_1")
	// check(err)
	// fmt.Print(string(dat))

	// // Output: [a b c]

}
