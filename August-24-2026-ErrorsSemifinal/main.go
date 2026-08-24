package main

import (
	"errors"
	"fmt"
)

// Эта функция делит x на y и возвращает ошибку, если y равен нулю
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	// Тестовые значения
	numerator := 10.0
	denominator := 0.0
	
	// Вызов функции divide
	result, err := divide(numerator, denominator)
	
	// TODO: Проверить, не является ли err отличным от nil (есть ли ошибка)
	// Если есть ошибка, вывести её с помощью fmt.Println(err)
	// В противном случае вывести результат с помощью fmt.Printf("Result: %.2f\n", result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
