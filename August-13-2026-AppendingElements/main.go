package main

import "fmt"

func main() {
	// Срез фруктов
	fruits := []string{"apple", "banana", "orange"}
	
	// TODO: Добавьте "grape" и "kiwi" в срез fruits
	// Напишите свой код здесь
	fruits = append(fruits, "grape", "kiwi")
	// Вывести обновленный срез
	fmt.Println("My fruit collection:", fruits)
}
