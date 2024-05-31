package main

type Link string

type KnowledgeItem interface {
	Filename() string
	Title() string
	Links() []Link
	Content() string
}

type TextDocument struct {
	Filename string
	Title    string
	Links    []Link
	Content  string
}

func NewTextDocument(filename string, title string, links []Link, content string) TextDocument {
	return TextDocument{
		Filename: filename,
		Title:    title,
		Links:    links,
		Content:  content,
	}
}
