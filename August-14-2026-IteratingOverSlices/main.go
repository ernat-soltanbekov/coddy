package main

import "fmt"

func main() {
	// Срез фруктов
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}
	
	// TODO: Завершите цикл for, используя range для итерации по срезу fruits
	// Выведите каждый фрукт с его позицией, см. инструкции для точного формата
	for i, value := range fruits {
		fmt.Printf("%d. %s\n", i + 1, value)
	}
}
