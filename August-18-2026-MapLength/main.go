package main

import "fmt"

func main() {
	// Карта стран и их столиц
	countries := map[string]string{
		"France":  "Paris",
		"Japan":   "Tokyo",
		"Brazil":  "Brasília",
		"Canada":  "Ottawa",
		"Egypt":   "Cairo",
	}
	
	// TODO: Используйте функцию len(), чтобы получить количество стран в карте
	// и сохраните его в переменной с именем countryCount
	countryCount := len(countries)
	
	// Вывести количество стран
	fmt.Printf("The map contains %d countries\n", countryCount)
}
