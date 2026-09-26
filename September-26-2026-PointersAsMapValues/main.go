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
	
	// TODO: Создать map для хранения указателей на продукты
	
	// TODO: Создать slice для поддержания порядка продуктов (чтобы обеспечить согласованный вывод)
	
	// TODO: Разобрать данные о продуктах и заполнить map
	// Не забудьте:
	// - Разделить productDataStr по запятым, чтобы получить отдельные записи
	// - Для каждой записи разделить по двоеточиям, чтобы получить name, price, quantity
	// - Преобразовать строки price и quantity в соответствующие типы
	// - Сохранить указатель на структуру Product в map
	// - Добавить имя продукта в order slice
	
	// TODO: Display initial inventory
	// Использовать order slice для последовательного перебора продуктов
	// Формат: "[name]: $[price] (Stock: [quantity])"
	
	// TODO: Разобрать и применить операции обновления
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
