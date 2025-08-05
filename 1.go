package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) GetInfo() string {
	return fmt.Sprintf("Title: %s, Author: %s, Year: %d\n", b.Title, b.Author, b.Year)
}
func main() {
	book1 := Book{
		Title:  "Руслан и Людмила",
		Author: "А.C. Пушкин",
		Year:   192,
	}
	fmt.Println(book1.GetInfo())
}
