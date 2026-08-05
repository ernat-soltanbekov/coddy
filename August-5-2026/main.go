package main

import "fmt"

// Эта функция должна возвращать три значения: name, age и isStudent
func getPersonInfo() (string, int, bool) {
	// Переменные уже определены для вас
	name := "Alex"
	age := 25
	isStudent := true
	
	// TODO: Верните все три значения (name, age, isStudent)
	return name, age, isStudent
}

func main() {
	// Вызовите функцию и сохраните возвращаемые значения
	name, age, isStudent := getPersonInfo()
	
	// Выведите значения
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
}
