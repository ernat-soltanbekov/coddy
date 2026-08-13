package main

import "fmt"

func main() {
	// Срез с 3 элементами, но емкостью 5
	numbers := make([]int, 3, 5)
	
	// TODO: Выведите длину среза с помощью len()
	fmt.Println("Length:", len(numbers))
	
	// TODO: Выведите емкость среза с помощью cap()
	fmt.Println("Capacity:", cap(numbers))
	
}
