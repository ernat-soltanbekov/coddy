package main

import "fmt"

func main() {
	// Карта оценок студентов
	grades := map[string]int{
		"John":  85,
		"Sarah": 92,
		"Mike":  78,
		"Lisa":  95,
	}

	// Студент для проверки
	studentToCheck := "Emma"

	// TODO: Проверьте, существует ли studentToCheck в карте grades
	// и сохраните результат в переменной с именем 'exists'
	// Подсказка: используйте идиому comma-ok
	_, exists := grades[studentToCheck]

	// Вывод результата
	if exists {
		fmt.Printf("%s's grade exists in the map\n", studentToCheck)
	} else {
		fmt.Printf("%s's grade does not exist in the map\n", studentToCheck)
	}
}
