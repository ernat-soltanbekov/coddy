package main

import "fmt"

// Структура Address содержит информацию о местоположении
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Структура Person должна включать в себя структуру Address
type Person struct {
	Name string
	Age  int
	Address
	// TODO: Вставьте структуру Address здесь (всего одна строка)
}

func main() {
	// Создайте нового человека с информацией об адресе
	person := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Wonderland",
			ZipCode: "12345",
		},
	}

	// Выведите информацию о человеке, включая адрес
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	
	// TODO: Выведите поля адреса напрямую из экземпляра person
	// (без использования person.Address.Street и т. д.)
	fmt.Println("Street:", person.Street)
	fmt.Println("City:", person.City)
	fmt.Println("ZipCode:", person.ZipCode)
}
