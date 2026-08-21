package main

import (
	"errors"
	"fmt"
)

func validateUsername(username string) error {
	// TODO: Проверить, если длина имени пользователя меньше 3 символов
	// Если это так, вернуть ошибку с сообщением "username too short"
	// В противном случае вернуть nil (нет ошибки)
	if len(username) < 3 {
		return errors.New("username too short")
	}
	return nil
}

func main() {
	// Тест с валидным именем пользователя
	validName := "bob123"
	err1 := validateUsername(validName)
	if err1 != nil {
		fmt.Printf("%s is invalid: %v\n", validName, err1)
	} else {
		fmt.Printf("%s is valid!\n", validName)
	}
	
	// Тест с невалидным именем пользователя
	invalidName := "ab"
	err2 := validateUsername(invalidName)
	if err2 != nil {
		fmt.Printf("%s is invalid: %v\n", invalidName, err2)
	} else {
		fmt.Printf("%s is valid!\n", invalidName)
	}
}
