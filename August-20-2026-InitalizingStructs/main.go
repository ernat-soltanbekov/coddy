package main

import "fmt"

// Определение структуры Person
type Person struct {
	Name       string
	Age        int
	IsEmployed bool
}

func main() {
	// TODO: Инициализируйте структуру Person с именем "Alice", возрастом 28 и isEmployed true
	// Используйте либо имена полей, либо синтаксис литерала структуры
	var alice Person

	alice = Person{Name: "Alice", Age: 28, IsEmployed: true}
	
	// Вывод информации о человеке
	fmt.Printf("Name: %s\n", alice.Name)
	fmt.Printf("Age: %d\n", alice.Age)
	fmt.Printf("Employed: %t\n", alice.IsEmployed)
}
