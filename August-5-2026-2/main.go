package main

import "fmt"

// Эта функция использует именованные возвращаемые значения
func getPersonInfo() (name string, age int, isStudent bool) {
	// TODO: Присвойте значения именованным возвращаемым переменным
	// name должно быть "Alex"
	// age должно быть 25
	// isStudent должно быть true
	name = "Alex"
	age = 25
	isStudent = true
	// Оператор "return" без аргументов вернет именованные возвращаемые значения
	return
}

func main() {
	name, age, isStudent := getPersonInfo()
	fmt.Printf("Name: %s\nAge: %d\nStudent: %t\n", name, age, isStudent)
}
