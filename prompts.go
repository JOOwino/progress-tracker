package main

type PromptFields struct {
	Options      []string
	Headline     string
	ChoicePrefix string
}

func AddBookPrompt() PromptFields {
	return PromptFields{
		Options: []string{
			"Title",
			"Chapters",
			"Pages",
			"Author",
		},
		Headline: "Enter Book Details",
	}
}

func ExistingFilePrompt() PromptFields {
	return PromptFields{
		Options: []string{
			"Add New Book",
			"Update Reading Progress",
			"List All Books",
			"Print Report",
			"Exit",
		},
		Headline:     "Choose an option:",
		ChoicePrefix: "Select a choice: ",
	}
}

func ListBooks(bookTitles []string) PromptFields {

	return PromptFields{
		Options:      bookTitles,
		Headline:     "Your books",
		ChoicePrefix: "Select resource to update: ",
	}

}

func NewFilePrompt() PromptFields {
	return PromptFields{
		Options: []string{
			"Add New Book",
			"Exit",
		},
		Headline:     "Choose an option:",
		ChoicePrefix: "Select a choice: ",
	}
}
