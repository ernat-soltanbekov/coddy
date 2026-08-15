package main

import "fmt"

func main() {
	// TODO: Создайте литерал карты с именем 'capitals', который сопоставляет страны с их столицами
	// Включите как минимум эти три пары:
	// "France" -> "Paris"
	// "Japan" -> "Tokyo"
	// "Brazil" -> "Brasilia"
	capitals := map[string]string {
		"France": "Paris",
		"Japan": "Tokyo",
		"Brazil": "Brasilia",
	}
	
	// Это выведет карту
	fmt.Println(capitals)
	
	// Это выведет столицу Японии
	fmt.Println("The capital of Japan is:", capitals["Japan"])
}
