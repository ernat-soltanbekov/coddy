package main

import (
	"fmt"
	"strconv"
)

func processNotification(value interface{}) {
	switch znachenie := value.(type) {
		case string:
		fmt.Printf("Email notification: %v", znachenie)
		case int:
		fmt.Printf("SMS notification with %v characters", znachenie)
		case bool:
		fmt.Printf("Push notifications: %v", znachenie)
		case float64:
		fmt.Printf("Alert with priority: %v", znachenie)
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
	switch notificationType {
	case "email":
		processNotification(content)
	case "sms":
		value, _ := strconv.Atoi(content)
		processNotification(value)
	case "push":
		value := content == "enabled"
		processNotification(value)
	case "alert":
		value, _ := strconv.ParseFloat(content, 64)
		processNotification(value)
	}
}
