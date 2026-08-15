package main

import "fmt"

func main() {
	// TODO: Создайте карту с именем 'favoriteColors' с помощью make
	// Ключи должны быть строками (именами), а значения также должны быть строками (цветами)
	favoriteColors := make(map[string]string)
	// Добавление нескольких пар ключ-значение в карту
	favoriteColors["Alice"] = "Blue"
	favoriteColors["Bob"] = "Green"
	favoriteColors["Charlie"] = "Red"
	
	// Вывод карты
	fmt.Println("Favorite colors:", favoriteColors)
	
	// Вывод любимого цвета Боба
	fmt.Println("Bob's favorite color is", favoriteColors["Bob"])
}
