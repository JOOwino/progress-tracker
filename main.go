package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
//In our scenario we work with Json File

func (rt *ReadTracker) LoadBooks() error {
	file, err := os.Open(rt.File)
	if err != nil {
		//If error is file does not exist://Create file
		if os.IsExist(err) {
			return generateNewFile(rt.File)
		}
		fmt.Printf("Another error, not of missing file: %v", err)
		return err

	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&rt.Books)

}

func generateNewFile(fileName string) error {
	_, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return err
	}
	return nil

}

func init() {
	fmt.Println("Welcome to Progress Tracker.A CLI tool to track Books read.")
	fmt.Println("Books are not Limited to Upskilling, but Fiction,Non-fiction And Biography")
}

func main() {
	options := []string{
		"Add New Book",
		"Update Reading Progress",
		"List All Books",
		"Print Report",
		"Exit",
	}

	choice := prompt(options, "Choose an option")
	fmt.Println("c: %d\n", choice)
}

func prompt(options []string, prompt string) int {

	for i, opt := range options {
		fmt.Printf("[%d]%s\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Select a choice: ")
		input, _ := reader.ReadString('\n')
		text := strings.TrimSpace(input)
		index, err := strconv.Atoi(text)
		//This assumes the user either entered wrong value(not a digit) || entered value greater than the options []string
		if err == nil && index > 0 && index <= len(options) {
			return index - 1 //converts to zero based index ?????
		}
		fmt.Println("Invalid choice. Enter a valid value")

	}

}
