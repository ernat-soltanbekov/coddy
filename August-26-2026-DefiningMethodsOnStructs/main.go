package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

// TODO: Определите вашу структуру Book здесь
type Book struct {
    Title string
    Author string
    Pages int
}

// TODO: Определите ваш метод displayInfo здесь
func (b Book) displayInfo() {
    fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.Title, b.Author, b.Pages)
}
// TODO: Определите ваш метод getDescription здесь
func (b Book) getDescription() {
    fmt.Println(b.Title, "by", b.Author)
}

func main() {
    // Создание сканера для чтения строк ввода
    scanner := bufio.NewScanner(os.Stdin)
    
    // Чтение названия
    scanner.Scan()
    title := scanner.Text()
    
    // Чтение автора
    scanner.Scan()
    author := scanner.Text()
    
    // Чтение количества страниц
    scanner.Scan()
    pagesStr := scanner.Text()
    
    // Преобразование строки страниц в целое число
    pages, _ := strconv.Atoi(pagesStr)

    book := Book {
        Title: title,
        Author: author,
        Pages: pages,
    }
    
    // TODO: Создайте экземпляр Book и вызовите методы
    book.displayInfo()
    book.getDescription()


}
