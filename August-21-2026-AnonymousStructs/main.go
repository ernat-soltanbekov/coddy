package main

import "fmt"

func main() {
	// TODO: Создайте переменную анонимной структуры с именем 'book' со следующими полями:
	// - title (string): "The Go Programming Language"
	// - pages (int): 380
	book := struct {
		title string
		pages int
	}{
		title: "The Go Programming Language",
		pages: 380,
	}
	
	// Выведите информацию о книге
	fmt.Printf("Book: %s, Pages: %d\n", book.title, book.pages)
}
