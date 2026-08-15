package main

import "fmt"

// addItems добавляет новые товары и их цены в список покупок
func addItems(names []string, prices []float64, newNames []string, newPrices []float64) ([]string, []float64) {
	names = append(names, newNames...)
    prices = append(prices, newPrices...)
	return names, prices
}

// removeItem удаляет товар по указанному индексу
func removeItem(names []string, prices []float64, index int) ([]string, []float64) {
	names = append(names[:index], names[index+1:]...) 
    prices = append(prices[:index], prices[index+1:]...)
	return names, prices
}

// findExpensiveItems возвращает товары с ценами выше порогового значения
func findExpensiveItems(names []string, prices []float64, threshold float64) []string {
	var result []string
	for i := range prices {
		if prices[i] > threshold {
			result = append(result, names[i])
		}
	}
	return result
}

// calculateTotalCost возвращает сумму всех цен
func calculateTotalCost(prices []float64) float64 {
    var sum float64
	for i := range prices {
        sum += prices[i]
    }
	return sum
}

func main() {
	// Инициализация пустых списков покупок
	names := []string{}
	prices := []float64{}
	
	// Добавление начальных товаров
	initialNames := []string{"Apples", "Milk", "Bread"}
	initialPrices := []float64{2.99, 3.49, 2.29}
	
	names, prices = addItems(names, prices, initialNames, initialPrices)
	
	// Печать начального списка покупок
	fmt.Println("Initial Shopping List:")
	for i := range names {
		fmt.Printf("%d. %s - $%.2f\n", i, names[i], prices[i])
	}
	
	// Вычисление и печать общей стоимости
	total := calculateTotalCost(prices)
	fmt.Printf("\nTotal Cost: $%.2f\n", total)
	
	// Поиск и печать дорогих товаров
	priceThreshold := 3.00
	expensiveItems := findExpensiveItems(names, prices, priceThreshold)
	
	fmt.Printf("\nExpensive Items (above $%.2f):\n", priceThreshold)
	for _, item := range expensiveItems {
		fmt.Println(item)
	}
	
	// Добавление нового товара
	names, prices = addItems(names, prices, []string{"Cheese"}, []float64{4.99})
	
	// Удаление товара
	names, prices = removeItem(names, prices, 1)
	fmt.Println("\nRemoved item at index 1")
	
	// Печать итогового списка покупок
	fmt.Println("\nFinal Shopping List:")
	for i := range names {
		fmt.Printf("%d. %s - $%.2f\n", i, names[i], prices[i])
	}
}
