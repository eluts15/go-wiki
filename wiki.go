package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("Sick test page bro. Love me.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
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
// Functions can return multiple values, and just toss out the uneeded value(through use of the blank identifer).
// In this case if it returns nil, than it was successful.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
