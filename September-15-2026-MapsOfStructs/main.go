package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
	"sort"
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
	divide := strings.Split(studentData, ",")
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
	names := make([]string, 0, len(student))
	for n := range student {
		names = append(names, n)
	}
	sort.Strings(names)
	for n := range names {
		fmt.Printf("%s: ID %d, Grade %s\n", names[n], student[names[n]].ID, student[names[n]].Grade)
	}
	// TODO: Рассчитайте и выведите статистику оценок
	grades := map[string]int{}
	for i := 0; i < len(names); i++ {
		grades[student[names[i]].Grade]++
	}
    	gradeOrder := []string{"A", "B", "C", "D", "F"}
    for i := 0; i < len(gradeOrder); i ++ {
        if grades[gradeOrder[i]] > 0 {
            fmt.Printf("Grade %s: %d students\n", gradeOrder[i], grades[gradeOrder[i]])
        }
    }
	// TODO: Найдите и выведите студента с самым высоким ID
    maxID := student[names[0]].ID
    maxIDname := names[0]
    for count := 0; count < len(names); count++ {
        if student[names[count]].ID > maxID {
            maxID = student[names[count]].ID
            maxIDname = names[count]
        }
    }
	// TODO: Выведите общее количество студентов
}
