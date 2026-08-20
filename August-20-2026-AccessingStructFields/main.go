package main

import "fmt"

func main() {
	// Структура Person с тремя полями
	type Person struct {
		name      string
		age       int
		isStudent bool
	}

	// Пример человека
	john := Person{
		name:      "John Doe",
		age:       25,
		isStudent: true,
	}

	// TODO: Выведите имя john, используя fmt.Printf и формат %s
	fmt.Printf("Name: %s\n", john.name)
	
	// TODO: Выведите возраст john, используя fmt.Printf и формат %d
	fmt.Printf("Age: %d\n", john.age)
	
	// TODO: Выведите, является ли john студентом, используя fmt.Printf и формат %t
	fmt.Printf("Is student: %t\n", john.isStudent)
	
}
