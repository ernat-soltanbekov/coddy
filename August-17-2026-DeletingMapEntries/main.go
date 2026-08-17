package main

import "fmt"

func main() {
	// Карта предметов инвентаря и их количества
	inventory := map[string]int{
		"pen":    15,
		"pencil": 10,
		"paper":  25,
		"eraser": 5,
	}

	fmt.Println("Initial inventory:", inventory)

	// TODO: Удалите запись 'pencil' из карты inventory
	delete(inventory, "pencil")

	fmt.Println("Updated inventory:", inventory)

	// Проверьте, существует ли еще 'pencil' в карте
	_, exists := inventory["pencil"]
	fmt.Println("Does 'pencil' exist in inventory?", exists)
}
