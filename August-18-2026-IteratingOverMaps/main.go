package main

import (
	"fmt"
	"sort"
)

func main() {
	// Карта оценок студентов
	studentGrades := map[string]string{
		"Alice":  "A",
		"Bob":    "B",
		"Charlie": "B+",
		"David":  "A-",
		"Eva":    "C",
	}

	// Собираем имена студентов (ключи) в срез, чтобы отсортировать их
	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	sort.Strings(names)


	// TODO: Пройтись по отсортированному списку имен студентов
	// и вывести имя и оценку каждого студента в формате:
	// "Student: [name], Grade: [grade]"
	for _, name := range names {
		fmt.Printf("Student: %s, Grade: %s\n", name, studentGrades[name])
	}
}
