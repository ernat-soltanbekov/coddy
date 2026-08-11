package main

import "fmt"

func main() {
	// Это наш массив любимых фруктов
	favoriteFruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Strawberry"}
	
	// TODO: Используйте функцию len(), чтобы найти длину массива favoriteFruits
	// и сохраните её в переменной numberOfFruits
	var numberOfFruits int
	numberOfFruits = len(favoriteFruits)
	
	// Это выведет количество фруктов в массиве
	fmt.Printf("There are %d fruits in the array\n", numberOfFruits)
}
