package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Прочитать ввод
	var numProductsStr string
	var productDataStr string
	var operationsStr string
	
	fmt.Scanln(&numProductsStr)
	fmt.Scanln(&productDataStr)
	fmt.Scanln(&operationsStr)
	
	// TODO: Определить структуру Product здесь
	type Product struct {
		Price float64
		Quantity int
	}
	// TODO: Создать map для хранения указателей на продукты
	product := map[string]*Product{}
	// TODO: Создать slice для поддержания порядка продуктов (чтобы обеспечить согласованный вывод)
	sliceOrder := []string{}
	// TODO: Разобрать данные о продуктах и заполнить map
	productData := strings.Split(productDataStr, ",")
	for i := 0; i < len(productData); i++ {
		productDescribe := strings.Split(productData[i], ":")
		price, err := strconv.ParseFloat(productDescribe[1], 64)
		if err != nil {
			fmt.Printf("Invalid price: %v\n", err)
			return
		}
		quantity, err := strconv.Atoi(productDescribe[2])
		if err != nil {
			fmt.Printf("Invalid quantity: %v\n", err)
			return
		}
		dataStructure := Product{price, quantity}
		product[productDescribe[0]] = &dataStructure
        sliceOrder = append(sliceOrder, product[productDescribe[0]])
	}
	// Не забудьте:
	// - Разделить productDataStr по запятым, чтобы получить отдельные записи
	// - Для каждой записи разделить по двоеточиям, чтобы получить name, price, quantity
	// - Преобразовать строки price и quantity в соответствующие типы
	// - Сохранить указатель на структуру Product в map
	// - Добавить имя продукта в order slice

	// TODO: Display initial inventory
    for i := 0; i < len(sliceOrder); i++ {
        fmt.Printf("Initial Inventory:\n")
        displayOrder := sliceOrder[i]
        price := product[displayOrder].Price
        quantity := product[displayOrder].Quantity
        fmt.Printf("%s: $%.2f (Stock: %d)\n", displayOrder, price, quantity)
    }
	// Использовать order slice для последовательного перебора продуктов
	// Формат: "[name]: $[price] (Stock: [quantity])"
	
	// TODO: Разобрать и применить операции обновления
    operation := strings.Split(operationsStr, ",")
	for i := 0; i < len(operation); i++ {
		operDescribe := strings.Split(operation[i], ":")
		opername, err := strconv.ParseFloat(operDescribe[2], 64)
		if err != nil {
			fmt.Printf("Invalid price: %v\n", err)
			return
		}
		if operDescribe[0] == "price" {
            product[operDescribe[1]].Price = opername
        }
	}
	// Не забудьте:
	// - Разделить operationsStr по запятым, чтобы получить отдельные операции
	// - Для каждой операции разделить по двоеточиям, чтобы получить type, name, value
	// - Обновить соответствующее поле напрямую через указатель map
	// - Вывести сообщение об обновлении для каждой операции
	
	// TODO: Display updated inventory
	// Снова используйте срез order для согласованного вывода
	// Вычисляйте общую стоимость во время отображения
	
	// TODO: Calculate and display total inventory value
	// Format: "Total Inventory Value: $[total_value]"
}
