package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Author            string `json:"author"`
	TotalPages        int    `json:"total_pages"`
	CurrentPage       int    `json:"current_page"`
	PercentageCovered string `json:"percentage_covered"`
	Status            string `json:"status"`
}

type ReadTracker struct {
	Books []Book `json:"books"`
	File  string `json:"-"`
}

func ReadingTracker(fileName string) *ReadTracker {
	return &ReadTracker{
		Books: []Book{},
		File:  fileName,
	}
}

//LoadBooks: Create an array of books from a file .
//If file is not present it creates one in the same directory.

func (rt *ReadTracker) LoadBooks() error {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 || text == "" {
			fmt.Printf("Error parsing input: %s \n", text)
		}
		fmt.Printf("Text parsed is: %s \n", text)

	}

}
