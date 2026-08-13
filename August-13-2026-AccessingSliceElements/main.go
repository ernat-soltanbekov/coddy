package main

import "fmt"

func main() {
	// Срез фруктов
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
	
	// TODO: Получите доступ к первому фрукту (индекс 0) и сохраните его в переменной с именем firstFruit
	firstFruit := fruits[0]
	// TODO: Получите доступ к третьему фрукту (индекс 2) и сохраните его в переменной с именем thirdFruit
	thirdFruit := fruits[2]
	// Выведите результаты
	fmt.Println("The first fruit is:", firstFruit)
	fmt.Println("The third fruit is:", thirdFruit)
}
