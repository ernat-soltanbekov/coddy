package main

import (
	"fmt"
)

func main() {
	// Предопределенное имя пользователя, которое слишком короткое
	username := "bob"
	// Минимально требуемая длина
	minLength := 5
	
	// Проверка, является ли имя пользователя допустимым
	if len(username) < minLength {
		// TODO: Создайте отформатированную ошибку с помощью fmt.Errorf, которая включает:
		// 1. Недопустимое имя пользователя
		// 2. Минимально требуемую длину
		// Пример формата: "invalid username: [username] is too short, minimum length is [minLength]"
		err := fmt.Errorf("invalid username: %s is too short, minimum length is %d", username, minLength) // Ваш код здесь
		
		// Вывод ошибки
		fmt.Println(err)
	} else {
		fmt.Println("Username is valid!")
	}
}
