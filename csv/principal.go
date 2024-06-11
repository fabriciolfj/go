package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func importBooksFromCSV(filename string) ([]Book, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var books []Book

	for _, record := range records {
		pages, err := strconv.Atoi(record[2])

		if err != nil {
			continue
		}

		books = append(books, Book{
			Title:  record[0],
			Author: record[1],
			Pages:  pages,
		})
	}

	return books, nil
}

func exportBooksToCSV(filename string, books []Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, book := range books {
		record := []string{book.Title, book.Author, strconv.Itoa(book.Pages)}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fileName := "books.csv"

	books := []Book{
		{
			Title:  "java",
			Author: "fabricio",
			Pages:  100,
		},
		{
			Title:  "teste",
			Author: "teste",
			Pages:  22,
		},
	}

	if err := exportBooksToCSV(fileName, books); err != nil {
		fmt.Printf("failed to export books to csv: %s\n", err)
	}

	importedBook, err := importBooksFromCSV(fileName)

	if err != nil {
		fmt.Printf("failed to import books from csv :%s\n", err)
	} else {
		fmt.Printf("imported books", importedBook)
	}
}
