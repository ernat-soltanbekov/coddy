package main

import "fmt"

// Структура Person с полями name и age
type Person struct {
	name string
	age  int
}

func main() {
	// Создание структуры Person
	person := Person{
		name: "John",
		age:  25,
	}

	// Создание указателя на структуру Person
	personPtr := &person

	// Вывод исходной информации о человеке
	fmt.Println("Original person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	// TODO: Используйте указатель, чтобы изменить имя человека на "Jane" и возраст на 30
	// Подсказка: Вы можете получить доступ к полям структуры через указатель, используя точечную нотацию
	personPtr.name = "Jane"
	personPtr.age = 30

	// Вывод обновленной информации о человеке
	fmt.Println("Updated person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}
