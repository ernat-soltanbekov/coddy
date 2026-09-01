package main

import (
	"fmt"
	"strconv"
)

// TODO: Определите интерфейс Shape с методом Area() float64
type Shape interface {
	Area() float64
}
// TODO: Определите структуру Circle с полем Radius (float64)
type Circle struct {
	Radius float64
}
// TODO: Определите структуру Rectangle с полями Width и Height (float64)
type Rectangle struct {
	Width float64
	Height float64
}
// TODO: Реализуйте метод Area() для Circle (используйте 3.14159 для pi)
func (a Circle) Area() float64 {
	return 3.14159 * a.Radius * a.Radius
}
// TODO: Реализуйте метод Area() для Rectangle
func (b Rectangle) Area() float64 {
	return b.Width * b.Height
}
// TODO: Определите функцию printShapeInfo, которая принимает Shape и выводит её площадь
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %g", s.Area())
}
func main() {
	// Прочитать ввод
	var shapeType string
	var dimension1 string
	var dimension2 string
	fmt.Scanln(&shapeType)
	fmt.Scanln(&dimension1)
	fmt.Scanln(&dimension2)

	// Преобразовать строковые входные данные в float64
	dim1, _ := strconv.ParseFloat(dimension1, 64)
	dim2, _ := strconv.ParseFloat(dimension2, 64)

	// TODO: Объявите переменную Shape
	var peremennaya Shape
	switch shapeType {
	// TODO: Если shapeType равен "circle", присвойте Circle, используя dim1 как Radius
	case "circle":
		peremennaya = Circle{Radius: dim1}
	// TODO: Если shapeType равен "rectangle", присвойте Rectangle, используя dim1 как Width и dim2 как Height
	case "rectangle":
		peremennaya = Rectangle{Width: dim1, Height: dim2}
	}
	// TODO: Вызовите printShapeInfo с созданной фигурой
	printShapeInfo(peremennaya)
}
