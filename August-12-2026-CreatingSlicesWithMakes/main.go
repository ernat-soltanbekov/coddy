package main

import "fmt"

func main() {
	// TODO: Создайте срез строк длиной 3 и емкостью 5, используя make
	// var names = ...
	names := make([]string, 3, 5)
	// Присвоение значений срезу
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"
	
	// Печать среза
	fmt.Println("Names:", names)
	
	// Печать длины и емкости среза
	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))
}
