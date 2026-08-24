package main

import (
	"errors"
	"fmt"
)

// celsiusToFahrenheit преобразует градусы Цельсия в градусы Фаренгейта
// Возвращает ошибку, если температура ниже абсолютного нуля (-273.15°C)
func celsiusToFahrenheit(celsius float64) (float64, error) {
	// TODO: Реализовать логику преобразования
	// 1. Проверить, не ниже ли температура абсолютного нуля (-273.15°C)
	// 2. Если значение допустимо, преобразовать по формуле: F = C × 9/5 + 32
	// 3. Вернуть соответствующее значение и ошибку
    f := celsius * 9/5 + 32
    if celsius < -273.15 {
        return 0, errors.New("Error: temperature below absolute zero")
    }
    return f, nil
}

func main() {
	// Тест допустимой температуры
	fmt.Println("Converting 25°C to Fahrenheit:")
	fahrenheit, err := celsiusToFahrenheit(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("25°C = %.2f°F\n", fahrenheit)
	}

	// Тест другой допустимой температуры
	fmt.Println("\nConverting 0°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("0°C = %.2f°F\n", fahrenheit)
	}

	// Тест недопустимой температуры (ниже абсолютного нуля)
	fmt.Println("\nConverting -300°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(-300)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("-300°C = %.2f°F\n", fahrenheit)
	}
}
