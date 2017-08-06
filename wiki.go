package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("vim-go")
}

// Describes how page data will be stored in memory.
type Page struct {
	Title string
	Body  []byte
}

// A function for persistent storage.
// of the type error, should anything go wrong, let Go handle it.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// A function for loading pages.
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
