package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func generateLibrarySummary(books []Book) {
	fmt.Printf("Total books: %d\n", len(books))

	var totalPages int
	booksByAuthor := make(map[string][]Book)
	for _, book := range books {
		totalPages += book.Pages
		booksByAuthor[book.Author] = append(booksByAuthor[book.Author], book)
	}

	avgPages := float64(totalPages) / float64(len(books))
	fmt.Printf("Average Pages per Book: %.2f\n", avgPages)

	for author, books := range booksByAuthor {
		fmt.Printf("%s has %d books\n", author, len(books))
	}
}

func exportLibraryDataForAnalysis(filename string, books []Book) error {
	file, err := os.Create(filename)

	if err != nil {
		return nil
	}

	defer file.Close()
	writer := csv.NewWriter(file)

	defer writer.Flush()

	if err := writer.Write([]string{"Title", "Author", "Pages"}); err != nil {
		return err
	}

	return nil
}

func main() {
	books := []Book{
		{
			Title:  "java",
			Author: "Fabricio",
			Pages:  10,
		},
		{
			Title:  "teste",
			Author: "Fabricio",
			Pages:  4,
		},
		{
			Title:  "teste",
			Author: "teste",
			Pages:  9,
		},
	}

	generateLibrarySummary(books)
	if err := exportLibraryDataForAnalysis("library.csv", books); err != nil {
		fmt.Println(err)
	}
}
