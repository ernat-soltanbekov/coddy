package main

import (
	"fmt"
	"strconv"
)

func processNotification(data interface{}) {
	switch value := data.(type) {
	case string:
		fmt.Printf("Email notification: %s\n", value)
	case int:
		fmt.Printf("SMS notification with %d characters\n", value)
	case bool:
		fmt.Printf("Push notifications: %t\n", value)
	case float64:
		fmt.Printf("Alert with priority: %g\n", value)
	default:
		fmt.Println("Unknown notification type")
	}
}

func main() {
	// Чтение входных данных
	var notificationType string
	var content string
	fmt.Scanln(&notificationType)
	fmt.Scanln(&content)
	
	// TODO: Напишите свой код ниже
	// 1. Создайте функцию processNotification, которая принимает параметр типа interface{}
	// 2. Преобразуйте content в соответствующий тип на основе notificationType
	// 3. Вызовите processNotification с преобразованным значением
	// 4. Используйте переключатель типов (type switch) внутри processNotification для обработки различных типов
	var data interface{}

	switch notificationType {
	case "email":
		data = content

	case "sms":
		value, _ := strconv.Atoi(content)
		data = value

	case "push":
		data = content == "enabled"

	case "alert":
		value, _ := strconv.ParseFloat(content, 64)
		data = value
	}

	processNotification(data)
}
