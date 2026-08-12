package main

import "fmt"

func main() {
	// TODO: Создайте литерал среза с именем 'colors', содержащий строки:
	// "red", "blue", "green" и "yellow"
	colors := []string{"red", "blue", "green", "yellow"}
	
	// Вывести срез
	fmt.Println("Colors:", colors)
	
	// Вывести длину среза
	fmt.Println("Number of colors:", len(colors))
}
