package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
    "maps"
)

func main() {
	// Чтение входных данных
	var numStudentsStr string
	var studentData string
	fmt.Scanln(&numStudentsStr)
	fmt.Scanln(&studentData)
	numbers, err = strconv.Atoi(numStudentsStr)
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
	sorted := maps.Keys(student)
    sort.Strings(sorted)
	// TODO: Рассчитайте и выведите статистику оценок
	
	// TODO: Найдите и выведите студента с самым высоким ID
	
	// TODO: Выведите общее количество студентов
}
