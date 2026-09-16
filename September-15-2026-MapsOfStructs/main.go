package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	// Чтение входных данных
	var numStudentsStr string
	var studentData string
	fmt.Scanln(&numStudentsStr)
	fmt.Scanln(&studentData)
	numbers, err := strconv.Atoi(numStudentsStr)
	if err != nil {
		fmt.Printf("Invalid number: %v", err)
		return
	}
	// TODO: Определите структуру Student здесь
	type Student struct {
		ID int
		Grade string
	}
	// TODO: Создайте карту (map) для хранения студентов (имя как ключ, структура Student как значение)
	student := map[string]Student{}
	// TODO: Разберите данные о студентах и заполните карту
	dividedStudentData := strings.Split(studentData, ",")
	for i := 0; i < len(dividedStudentData); i++ {
		doubleDividedStudentData := strings.Split(dividedStudentData[i], ":")
        numberID, err := strconv.Atoi(doubleDividedStudentData[1])
	    if err != nil {
	    	fmt.Printf("Invalid ID: %v\n", err)
	    	return
	    }
        students := Student{numberID, doubleDividedStudentData[2]}
        student[doubleDividedStudentData[0]] = students
	}
	// TODO: Выведите всех студентов в алфавитном порядке по имени
	sorted := []string{}
    for key := range student {
        sorted = append(sorted, key)
    }
    sort.Strings(sorted)
    for print := 0; print < len(sorted); print++ {
        fmt.Printf("%s: ID %d, Grade %s\n", sorted[print], student[sorted[print]].ID, student[sorted[print]].Grade)
    }
	// TODO: Рассчитайте и выведите статистику оценок
    gradeScore := map[string]int{}
	for gradeCount := 0; gradeCount < len(sorted); gradeCount++ {
        gradeScore[student[sorted[gradeCount]].Grade]++ 
    }
    gradeAlpha := []string{"A", "B", "C", "D", "F"}
    for count := 0; count < len(gradeAlpha); count++ {
        if gradeScore[gradeAlpha[count]] > 0 {
            fmt.Printf("Grade %s: %d students\n", gradeAlpha[count], gradeScore[gradeAlpha[count]])
        }
    }
	// TODO: Найдите и выведите студента с самым высоким ID
    maxID := student[sorted[0]].ID
    maxIDname := sorted[0]
    for i := 0; i < len(sorted); i++ {
        if student[sorted[i]].ID > maxID {
            maxID = student[sorted[i]].ID
            maxIDname = sorted[i]
        }
    }
	// TODO: Выведите общее количество студентов
    fmt.Printf("Highest ID: %s (%d)\n", maxIDname, maxID)
	// TODO: Выведите общее количество студентов
    fmt.Printf("Total students: %d\n", numbers)
}
