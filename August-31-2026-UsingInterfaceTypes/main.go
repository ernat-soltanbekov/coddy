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
func (r Circle) Area() float64 {
	return 3.14159 * r.Radius * r.Radius
}
// TODO: Реализуйте метод Area() для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
// TODO: Определите функцию printShapeInfo, которая принимает Shape и выводит её площадь
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %g\n", s.Area())
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
	var forma Shape
	// TODO: Если shapeType равен "circle", присвойте Circle, используя dim1 как Radius
	if shapeType == "circle" {
		forma = Circle{Radius: dim1}
	}
	// TODO: Если shapeType равен "rectangle", присвойте Rectangle, используя dim1 как Width и dim2 как Height
	if shapeType == "rectangle" {
		forma = Rectangle{Width: dim1, Height: dim2}
	}
	// TODO: Вызовите printShapeInfo с созданной фигурой
	printShapeInfo(forma)
}
