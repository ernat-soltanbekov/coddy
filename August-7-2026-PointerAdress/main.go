package main

import "fmt"

func main() {
	// Это наша переменная-образец
	name := "Gopher"
	
	// Это объявляет переменную-указатель для хранения адреса строки
	var namePointer *string
	
	// TODO: Используйте оператор взятия адреса (&), чтобы получить адрес памяти 'name'
	namePointer = &name
	// и сохраните его в переменной 'namePointer'
	
	// Вывести значение, на которое указывает указатель
	fmt.Printf("The value at that memory address is: %v\n", *namePointer)
}
