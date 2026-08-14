package main

import "fmt"

func main() {
	// Исходный слайс с названиями фруктов
	source := []string{"apple", "banana", "cherry", "date"}
	
	// Создаем целевой слайс емкостью 3 элемента
	destination := make([]string, 3)
	
	// TODO: Используйте функцию copy для копирования элементов из source в destination
	copy(destination, source)
	
	// Выводим целевой слайс
	fmt.Println("Destination slice:", destination)
}
