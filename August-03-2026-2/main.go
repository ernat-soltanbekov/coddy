package main

import "fmt"

// TODO: Добавьте параметры в эту функцию, чтобы она принимала приветствие и имя
// Функция должна принимать два строковых параметра: greeting и name
// Вы можете записать их как (greeting string, name string) или (greeting, name string)
func createGreeting(greeting, name string) string {
	// TODO: Верните строку, которая объединяет приветствие и имя
	return greeting + ", " + name + "!"
	// Например: "Hello, John!"
	// Используйте формат: greeting + ", " + name + "!"
}

func main() {
	// Эти тестовые вызовы уже определены за вас
	message := createGreeting("Hello", "Gopher")
	fmt.Println(message)
	
	// Тест с другими значениями
	message = createGreeting("Welcome", "Friend")
	fmt.Println(message)
}
