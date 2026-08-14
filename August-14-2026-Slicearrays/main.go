package main

import "fmt"

func main() {
	// Вот наш исходный срез фруктов
	fruits := []string{"apple", "banana", "orange", "grape", "kiwi"}
	
	// TODO: Создайте новый срез с именем 'firstThree', содержащий только первые три фрукта
	firstThree := fruits[:3]
	// TODO: Создайте новый срез с именем 'lastTwo', содержащий только последние два фрукта
	lastTwo := fruits[3:]
	// Вывод результатов
	fmt.Println("First three fruits:", firstThree)
	fmt.Println("Last two fruits:", lastTwo)
}
