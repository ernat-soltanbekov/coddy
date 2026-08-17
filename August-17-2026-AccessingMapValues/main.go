package main

import "fmt"

func main() {
	// Карта оценок студентов
	studentGrades := map[string]string{
		"John":  "B+",
		"Emma":  "A-",
		"Sarah": "A",
		"Mike":  "C",
	}
	
	// TODO: Получите и сохраните оценку Эммы в переменной с именем emmaGrade
	emmaGrade := studentGrades["Emma"]
	// TODO: Выведите оценку Эммы с сообщением вида: "Emma's grade is: A-"
	fmt.Println("Emma's grade is:", emmaGrade)
}
