package main

import (
	"fmt"                              // fmt нужен для чтения данных из терминала и вывода результата
	"strconv"                          // strconv нужен для преобразования строк в числа
)

type Rectangle struct {               // Создаём собственный тип Rectangle для хранения прямоугольника
	Width float64                     // Width хранит ширину прямоугольника
	Height float64                    // Height хранит высоту прямоугольника
}

func (a Rectangle) area() float64 {   // Value receiver: метод получает копию Rectangle и только читает его данные
	return a.Width * a.Height         // Площадь равна ширине, умноженной на высоту; результат возвращаем как float64
}

func (a *Rectangle) scale(scaleFactor float64) { // Pointer receiver: нужен указатель, потому что метод изменяет исходный Rectangle
	a.Width *= scaleFactor             // Умножаем ширину на коэффициент масштабирования и сохраняем новое значение
	a.Height *= scaleFactor            // Умножаем высоту на коэффициент масштабирования и сохраняем новое значение
}

func main() {                         // Точка входа программы
	var widthStr string               // Здесь ширина сначала хранится как строка, потому что ввод из терминала — текст
	var heightStr string              // Здесь высота сначала хранится как строка
	var scaleStr string                // Здесь коэффициент масштабирования сначала хранится как строка
	
	fmt.Scanln(&widthStr)              // Читаем ширину из терминала и записываем её в widthStr
	fmt.Scanln(&heightStr)             // Читаем высоту из терминала и записываем её в heightStr
	fmt.Scanln(&scaleStr)              // Читаем коэффициент масштабирования и записываем его в scaleStr
	
	width, _ := strconv.ParseFloat(widthStr, 64) // Превращаем строку ширины в число float64
	height, _ := strconv.ParseFloat(heightStr, 64) // Превращаем строку высоты в число float64
	scaleFactor, _ := strconv.ParseFloat(scaleStr, 64) // Превращаем строку коэффициента в число float64
	
	rectangle := Rectangle{            // Создаём экземпляр Rectangle
		Width: width,                    // Передаём введённую ширину в поле Width
		Height: height,                  // Передаём введённую высоту в поле Height
	}
	
	initialArea := rectangle.area()    // Вызываем area: метод берёт текущие Width и Height и возвращает их произведение
	fmt.Printf("Initial area: %.0f\n", initialArea) // Выводим начальную площадь; %.0f показывает число без знаков после точки
	
	rectangle.scale(scaleFactor)       // Передаём коэффициент в scale; pointer receiver изменяет сам rectangle
	
	scaledArea := rectangle.area()     // Снова вычисляем площадь, но теперь уже после изменения Width и Height
	fmt.Printf("Scaled area: %.0f\n", scaledArea) // Выводим площадь после масштабирования
}
