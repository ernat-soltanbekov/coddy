package main

import "fmt"

func main() {
	// Массив фруктов
	fruits := [5]string{"Apple", "Banana", "Orange", "Grape", "Mango"}
	
	// TODO: Завершите цикл for, чтобы вывести каждый фрукт из массива
	for i := 0; i < len(fruits); i++ {
		// Добавьте свой код здесь, чтобы вывести каждый фрукт
		fmt.Println(fruits[i])
	}
}
