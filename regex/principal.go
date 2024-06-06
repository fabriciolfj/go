package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var bookDetailsPattern = regexp.MustCompile(`(?s)Title:\[(.*?)\],\s*Author:\[(.*?)\],\s*Pages:\[(.*?)\]`)

type Book struct {
	Author string
	Title  string
	Pages  int
}

func ParseBooksFromFile(filename string) ([]Book, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var books []Book
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		fmt.Println(scanner.Text())

		if matches != nil && len(matches) == 4 {
			title := matches[1]
			author := matches[2]
			pages, err := strconv.Atoi(matches[3])

			if err != nil {
				fmt.Printf("invalid page number for book '%s\n'", title, err)
				continue
			}

			books = append(books, Book{Title: title, Author: author, Pages: pages})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func main() {
	filename := "books_listings.txt"

	books, err := ParseBooksFromFile(filename)

	if err != nil {
		fmt.Println("Error parsing books from file:", err)
		return
	}

	for _, book := range books {
		fmt.Printf("Parsed Book: %+v\n", book)
	}
}
