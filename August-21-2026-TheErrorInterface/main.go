package main

import (
	"errors"
	"fmt"
)

func main() {
	// Эта функция возвращает ошибку
	err := generateError()
	
	// TODO: Проверьте, что err не равен nil
	// Если err не равен nil, выведите: "Error occurred: " за которым следует сообщение об ошибке
	// HINT: Используйте оператор if для проверки, что err != nil
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	
	// Эта строка выполняется независимо от того, произошла ли ошибка
	fmt.Println("Program completed successfully")
}

// Эта функция генерирует ошибку
func generateError() error {
	return errors.New("something went wrong")
}
