package main

import (
	"errors"
	"fmt"
)

func main() {
	// Тест с допустимым именем пользователя
	result, err := validateUsername("gopher123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Тест с недопустимым именем пользователя
	result, err = validateUsername("go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func validateUsername(username string) (string, error) {
	if len(username) < 5 {
		// TODO: Вернуть ошибку с сообщением "username must be at least 5 characters long"
		// Подсказка: используйте errors.New() для создания новой ошибки
		return "", errors.New("username must be at least 5 characters long") // Замените эту строку
	}
	return "Username is valid", nil
}
