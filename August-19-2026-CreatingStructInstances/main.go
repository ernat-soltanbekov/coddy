package main

import "fmt"

// Определение структуры Person
type Person struct {
	name      string
	age       int
	isStudent bool
}

func main() {
	// TODO: Создайте новый экземпляр структуры Person с
	// name: "Alice", age: 25, isStudent: true
	person := Person{name: "Alice", age: 25, isStudent: true}
	// Не изменяйте код ниже
	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)
	fmt.Printf("Is Student: %t\n", person.isStudent)
}
