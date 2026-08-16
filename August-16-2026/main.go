package main

import (
    "fmt"
    "sort"
)

func main() {
	// Карта людей и их любимых фруктов
	favoriteFruits := map[string]string{
		"Alice": "Apple",
		"Bob":   "Banana",
		"Carol": "Cherry",
	}

	// TODO: Добавить новую запись для "David" с любимым фруктом "Dragon Fruit"
	favoriteFruits["David"] = "Dragon Fruit"
	// TODO: Обновить любимый фрукт Боба на "Blueberry"
	favoriteFruits["Bob"] = "Blueberry"
	// Собрать имена (ключи) в срез, чтобы отсортировать их
	var names []string
	for person := range favoriteFruits {
		names = append(names, person)
	}
	sort.Strings(names)

	// Вывести обновленную карту в отсортированном порядке
	fmt.Println("Updated favorite fruits:")
	for _, person := range names {
		fruit := favoriteFruits[person]
		fmt.Printf("%s loves %s\n", person, fruit)
	}
}
