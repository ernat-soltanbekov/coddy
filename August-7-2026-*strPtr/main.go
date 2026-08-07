package main

import (
	"fmt"
	"strings"
)

// makeUppercase принимает указатель на строку и переводит строку в верхний регистр
func makeUppercase(strPtr *string) {
	// TODO: Используйте strings.ToUpper(), чтобы изменить строку, на которую указывает strPtr
	// Подсказка: вам нужно разыменовать указатель, чтобы получить доступ к значению строки
	*strPtr = strings.ToUpper(*strPtr)
	
}

func main() {
	// У нас уже есть строковая переменная
	message := "hello, world"
	
	// Выводим оригинальное сообщение
	fmt.Printf("Original message: %s\n", message)
	
	// Вызываем makeUppercase, передавая адрес message
	makeUppercase(&message)
	
	// Выводим измененное сообщение
	fmt.Printf("Modified message: %s\n", message)
}
