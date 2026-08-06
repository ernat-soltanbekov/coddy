package main

import "fmt"

func main() {
	// Вот переменная и указатель на нее
	originalValue := 42
	pointerToValue := &originalValue
	secondPointer := pointerToValue
	
	// TODO: Объявите новую переменную-указатель с именем 'secondPointer' 
	// которая указывает на тот же адрес памяти, что и 'pointerToValue'
	
	
	// Не изменяйте код ниже
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
	
	// Давайте изменим исходное значение и увидим, что все указатели отражают это изменение
	originalValue = 100
	fmt.Printf("\nAfter changing original value to 100:\n")
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
}
