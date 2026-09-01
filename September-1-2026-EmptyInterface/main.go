package main

import (
	"fmt"
	"strconv"
	"strings"
)
// TODO: Напишите ваш код ниже
// 1. Создайте функцию processData, которая принимает interface{}
// 2. Преобразуйте valueStr в соответствующий тип на основе dataType
// 3. Вызовите processData с преобразованным значением
func processData(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	// Считать ввод
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)

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
	kusok := strings.Split(valueStr, ",")
	znacheniya := make([]int, 0, len(kusok))
	for _, kusochek := range kusok {
	znachenie, _ := strconv.Atoi(kusochek)
	znacheniya = append(znacheniya, znachenie)
	}
	processData(znacheniya)
	}
}
