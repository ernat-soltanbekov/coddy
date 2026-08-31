package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processData(data interface{}) {
	fmt.Printf("Value: %v, Type: %T", data, data)
	}

func main() {
	// Считать ввод
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)
	
	// TODO: Напишите ваш код ниже
	// 1. Создайте функцию processData, которая принимает interface{}
	// 2. Преобразуйте valueStr в соответствующий тип на основе dataType
	// 3. Вызовите processData с преобразованным значением
	switch dataType {
	case "int":
	value, _ := strconv.Atoi(valueStr)
	processData(value)

	case "string":
	processData(valueStr)

	case "bool":
	value, _ := strconv.ParseBool(valueStr)
	processData(value)

	case "slice":
	parts := strings.Split(valueStr, ",")
	values := make([]int, 0, len(parts))

	for _, part := range parts {
		value, _ := strconv.Atoi(part)
		values = append(values, value)
		}
		processData(values)
	}
}
