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

		fileLines := strings.Split(string(dat), "\n")
		// fmt.Println(fileLines[1])

		dirtyGenre := strings.Split(fileLines[2], " ")
		genre := (strings.Join(dirtyGenre[1:], " "))

		if err != nil {
			// ... handle error
			panic(err)
		}

		dirtyYear := strings.Split(fileLines[3], " ")
		year, err := strconv.Atoi(dirtyYear[1])
		if err != nil {
			// ... handle error
			panic(err)
		}

		dirtyRating := strings.Split(fileLines[4], " ")
		rating, err := strconv.Atoi(dirtyRating[1])
		if err != nil {
			// ... handle error
			panic(err)
		}

		var book = Book{Name: name, Author: fileLines[1], Genre: genre, Year: year, Rating: rating}
		books = append(books, book)
	}
	fmt.Println(books)

	// books := make([]Book, 0, 5)
	// dat, err := os.ReadFile("books/test_1")
	// check(err)
	// fmt.Print(string(dat))

	// // Output: [a b c]

}
