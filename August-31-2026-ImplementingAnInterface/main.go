package main

import (
	"fmt"
	"strconv"
)

// TODO: 1. Определите интерфейс Vehicle с методом GetInfo() string
type Vehicle interface {
	GetInfo() string
}
// TODO: 2. Определите структуру Car с полями Brand (string) и Speed (int)
type Car struct {
	Brand string
	Speed int
}
// TODO: 3. Реализуйте метод GetInfo() для Car
func (g Car) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Speed: %d km/h", g.Brand, g.Speed)
}
func main() {
	// Чтение входных данных
	var brand string
	var speedStr string
	fmt.Scanln(&brand)
	fmt.Scanln(&speedStr)

	// Преобразование строки скорости в целое число
	speed, _ := strconv.Atoi(speedStr)

	// TODO: 4. Создайте экземпляр Car, используя brand и speed, затем вызовите GetInfo()
	car := Car {
	Brand: brand,
	Speed: speed,
	}
	// Вывод результата
	fmt.Println(car.GetInfo())
}
