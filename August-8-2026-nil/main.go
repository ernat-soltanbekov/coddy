package main

import "fmt"

func main() {
	// Это nil-указатель на целое число
	var ptr *int
	
	// TODO: Проверить, является ли ptr nil
	if ptr == nil {
		fmt.Println("The pointer is nil")
	} else {
		fmt.Println("The pointer is nil")
	}
	// Если он равен nil, вывести "The pointer is nil"
	// Если он не равен nil, вывести "The pointer is not nil"
	
	fmt.Println("Program completed")
}
