package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Чтение входных данных
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)
	
	// Переменная для хранения значения интерфейса
	var interfaceValue interface{}
	
	// TODO: Напишите свой код ниже
	// 1. Преобразуйте valueStr в соответствующий тип на основе dataType и сохраните в interfaceValue
	switch dataType {
	case "int":
		value, _ := strconv.Atoi(valueStr)
		interfaceValue = value
	case "string":
		interfaceValue = valueStr
	case "bool":
		value, _ := strconv.ParseBool(valueStr)
		interfaceValue = value
	}
	// 2. Используйте утверждение типа (type assertion), чтобы проверить, содержит ли interfaceValue ожидаемый тип
	switch dataType {
	case "int":
		value, ok := interfaceValue.(int)
		if ok {
			fmt.Printf("Success: %v is a %T\n", value, interfaceValue)
		} else {
			fmt.Printf("Failed: value is not a %T\n", interfaceValue)
		}
	case "string":
		value, ok := interfaceValue.(string)
		if ok {
			fmt.Printf("Success: %v is a %T\n", value, interfaceValue)
		} else {
			fmt.Printf("Failed: value is not a %T\n", interfaceValue)
		}
	case "bool":
		value, ok := interfaceValue.(bool)
		if ok {
			fmt.Printf("Success: %v is a %T\n", value, interfaceValue)
		} else {
			fmt.Printf("Failed: value is not a %T\n", interfaceValue)
		}
	}
	// 3. Выведите соответствующее сообщение об успехе или ошибке
}
