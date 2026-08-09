package main

import "fmt"

// updateQuantity обновляет количество товара и возвращает новое количество
// и информацию о том, закончился ли товар (количество = 0)
func updateQuantity(quantity *int, change int) (int, bool) {
    // TODO: Обновить значение по указателю количества и вернуть новое значение и статус отсутствия на складе
    *quantity += change
    if *quantity < 0 {
        *quantity = 0
        }
    return *quantity, *quantity == 0
}

// calculateValue вычисляет общую стоимость товара на основе количества и цены
// Функция также обновляет mostValuableItem, если этот товар дороже
func calculateValue(itemName string, quantity int, price float64, mostValuableItem *string, highestValue *float64) float64 {
    // TODO: Вычислить общую стоимость, обновить mostValuableItem при необходимости и вернуть общую стоимость
    totalValue := float64(quantity) * price
    if totalValue > *highestValue {
        *highestValue = totalValue
        *mostValuableItem = itemName
    }
    return totalValue
}

// displayInventory выводит сведения об инвентаре
// Используются указатели, чтобы функция могла отображать данные в реальном времени
func displayInventory(apples, oranges, bananas *int, applePrice, orangePrice, bananaPrice float64) {
    fmt.Printf("Apples: %d (Value: $%.2f)\n", *apples, float64(*apples) * applePrice)
    fmt.Printf("Oranges: %d (Value: $%.2f)\n", *oranges, float64(*oranges) * orangePrice)
    fmt.Printf("Bananas: %d (Value: $%.2f)\n", *bananas, float64(*bananas) * bananaPrice)
}

func main() {
    // Инициализация инвентаря
    apples := 10
    oranges := 15
    bananas := 8
    
    // Цены
    applePrice := 0.5  // $0.50 за шт.
    orangePrice := 0.7 // $0.70 за шт.
    bananaPrice := 0.3 // $0.30 за шт.
    
    // Отслеживание самого дорогого товара
    var mostValuableItem string
    var highestValue float64
    
    // Отображение начального инвентаря
    fmt.Println("Initial Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Вычисление начальной стоимости и поиск самого дорогого товара
    appleValue := calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue := calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue := calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Симуляция продаж
    fmt.Println("Processing sales...")
    _, applesOutOfStock := updateQuantity(&apples, -4) // Продать 4 яблока
    _, orangesOutOfStock := updateQuantity(&oranges, -8) // Продать 8 апельсинов
    _, bananasOutOfStock := updateQuantity(&bananas, -10) // Попытка продать 10 бананов (больше, чем есть)
    
    // Проверка, закончились ли товары
    if applesOutOfStock {
        fmt.Println("Apples are out of stock!")
    }
    if orangesOutOfStock {
        fmt.Println("Oranges are out of stock!")
    }
    if bananasOutOfStock {
        fmt.Println("Bananas are out of stock!")
    }
    
    // Отображение обновленного инвентаря
    fmt.Println("\nUpdated Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Сброс отслеживания самого дорогого товара для перерасчета
    mostValuableItem = ""
    highestValue = 0
    
    // Пересчет стоимости
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n\n", mostValuableItem)
    
    // Пополнение запасов
    fmt.Println("Restocking...")
    updateQuantity(&apples, 5)  // Добавить 5 яблок
    updateQuantity(&oranges, 10) // Добавить 10 апельсинов
    updateQuantity(&bananas, 12) // Добавить 12 бананов
    
    // Отображение финального инвентаря
    fmt.Println("\nFinal Inventory:")
    displayInventory(&apples, &oranges, &bananas, applePrice, orangePrice, bananaPrice)
    
    // Сброс отслеживания самого дорогого товара для финального расчета
    mostValuableItem = ""
    highestValue = 0
    
    // Финальный расчет стоимости
    appleValue = calculateValue("Apples", apples, applePrice, &mostValuableItem, &highestValue)
    orangeValue = calculateValue("Oranges", oranges, orangePrice, &mostValuableItem, &highestValue)
    bananaValue = calculateValue("Bananas", bananas, bananaPrice, &mostValuableItem, &highestValue)
    
    fmt.Printf("Total inventory value: $%.2f\n", appleValue+orangeValue+bananaValue)
    fmt.Printf("Most valuable item: %s\n", mostValuableItem)
}
