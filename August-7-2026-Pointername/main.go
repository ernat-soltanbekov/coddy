package main

import "fmt"

func main() {
	// Вот наша строковая переменная
	message := "Hello, pointers!"
	
	// Это указатель на нашу переменную message
	messagePtr := &message
	
	// TODO: Разыменуйте messagePtr, чтобы получить значение, на которое он указывает,
	// и сохраните его в переменной 'value'
	value := string(*messagePtr)
	
	// Вывод результатов
	fmt.Println("The pointer points to the value:", value)
}
