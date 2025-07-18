package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var newFile bool
var prompt PromptFields

type Book struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	Author            string `json:"author"`
	Chapter           int    `json:"chapter"`
	TotalPages        int    `json:"total_pages"`
	CurrentPage       int    `json:"current_page"`
	PercentageCovered int    `json:"percentage_covered"`
	Status            string `json:"status"`
}

type ReadTracker struct {
	Books []Book `json:"books"`
	File  string `json:"-"`
}

func NewReadingTracker(fileName string) *ReadTracker {
	return &ReadTracker{
		Books: []Book{},
		File:  fileName,
	}
}

//LoadBooks: Create an array of books from a file .
//If file is not present it creates one in the same directory.
//In our scenario we work with Json File

func (rt *ReadTracker) LoadBooks() error {

	//Get Executable path so as to create file or read from same directory
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error fetching the path: %v\n", err)
	}

	dir := filepath.Dir(execPath)
	filePath := dir + "/" + rt.File
	rt.File = filePath

	_, err = os.Stat(filePath)
	if err != nil {
		fmt.Printf("PathError: File not found: %v\n", err)
		return generateNewFile(filePath)
	}
	file, err := os.Open(rt.File)
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
	ptracker := NewReadingTracker("tracker.json")
	err := ptracker.LoadBooks()
	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		if err == io.EOF {
			newFile = true
		} else {
			//No need to continue at this point. No file to load/read from || save to.
			os.Exit(9)
		}

	}
	if newFile {
		prompt = NewFilePrompt()
	} else {
		prompt = ExistingFilePrompt()
	}

	choice := multipleOptionsPrompt(prompt)

	switch choice {
	case 1:
		ptracker.AddBook()
	default:
		fmt.Println("Default Option")

	}
}

func (rt *ReadTracker) AddBook() {

	title := singleOptionPrompt("Book Title: ")
	author := singleOptionPrompt("Book Author: ")
	chapters := promptNumberConversion("Number of Chapters: ")
	pages := promptNumberConversion("Number of Pages: ")

	book := Book{
		ID:                generateID(),
		Title:             title,
		Author:            author,
		TotalPages:        pages,
		CurrentPage:       0,
		PercentageCovered: 0,
		Chapter:           chapters,
	}

	rt.Books = append(rt.Books, book)
	rt.saveFile()

}

func (rt *ReadTracker) saveFile() {
	file, err := os.OpenFile(rt.File, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(rt.Books); err != nil {
		fmt.Printf("Error encoding book data: %v\n", err)
	}

}

func multipleOptionsPrompt(promptData PromptFields) int {

	fmt.Println(promptData.Headline)
	for i, opt := range promptData.Options {
		fmt.Printf("[%d]%s\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(promptData.ChoicePrefix)
		input, _ := reader.ReadString('\n')
		text := strings.TrimSpace(input)
		index, err := strconv.Atoi(text)
		//This assumes the user either entered wrong value(not a digit) || entered value greater than the options []string
		if err == nil && index > 0 && index <= len(promptData.Options) {
			//If a valid value Clear Terminal and Return The value to preselect the next prompt
			clearTerminal()
			return index //converts to zero based index ?????
		}
		fmt.Println("Invalid choice. Enter a valid value")

	}

}

func promptNumberConversion(prompt string) int {
	value, err := strconv.Atoi(singleOptionPrompt(prompt))
	if err != nil {
		fmt.Println("The value should be a number")
		return promptNumberConversion(prompt)
	}
	return value
}

func singleOptionPrompt(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
