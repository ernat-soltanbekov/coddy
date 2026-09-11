package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

func main() {
	// Чтение входных данных
	var numStudentsStr string
	var studentData string
	fmt.Scanln(&numStudentsStr)
	fmt.Scanln(&studentData)
	
	// TODO: Определите структуру Student здесь
	type Student struct {
		ID int
		Grade string
	}
	// TODO: Создайте карту (map) для хранения студентов (имя как ключ, структура Student как значение)
	student := map[string]Student{}
	// TODO: Разберите данные о студентах и заполните карту
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//scanInput := scanner.Text()
	divide := strings.Split(scanInput, ",")
	for count := 0; count < len(divide); count++ {
		notes := strings.Split(divide[count], ":")
		notesDigit, err := strconv.Atoi(notes[1])
		if err != nil {
			fmt.Printf("Invalid digit: %v", err)
			return
		}
		noteStudent := Student{notesDigit, notes[2]}
		student[notes[0]] = noteStudent
	}
	// TODO: Выведите всех студентов в алфавитном порядке по имени
	
	// TODO: Рассчитайте и выведите статистику оценок
	
	// TODO: Найдите и выведите студента с самым высоким ID
	
	// TODO: Выведите общее количество студентов
}
