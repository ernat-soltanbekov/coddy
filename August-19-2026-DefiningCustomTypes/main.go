package main

import "fmt"

// TODO: Определите пользовательский тип с именем 'Temperature' на основе float64

func main() {
	// Создайте переменную типа Temperature со значением 23.5
	type Temperature float64
	var roomTemp Temperature = 23.5
	// Выведите значение и тип roomTemp
	fmt.Printf("Room temperature: %v\n", roomTemp)
	fmt.Printf("Type of roomTemp: %T\n", roomTemp)
}
