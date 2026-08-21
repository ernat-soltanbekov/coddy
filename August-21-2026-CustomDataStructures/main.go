package main

import (
	"fmt"
)

// Counter — это простая структура данных счетчика
type Counter struct {
	value int
	name  string
}

// NewCounter создает новый счетчик с заданным именем
func NewCounter(name string) Counter {
	return Counter{
		value: 0,
		name:  name,
	}
}

// Increment увеличивает счетчик на 1
func (c *Counter) Increment() {
	// Добавьте код для увеличения значения счетчика
    c.value++
}

// Decrement уменьшает счетчик на 1 (но не ниже 0)
func (c *Counter) Decrement() {
	// Добавьте код для уменьшения значения счетчика
    c.value--
	// Убедитесь, что значение не опускается ниже 0
    if c.value < 0 {
        c.value = 0
    }
}

// Reset сбрасывает счетчик обратно на 0
func (c *Counter) Reset() {
	// Добавьте код для сброса счетчика на 0
    c.value = 0
}

// Value возвращает текущее значение счетчика
func (c Counter) Value() int {
	// Добавьте код для возврата текущего значения
	return c.value // Замените это // Уже заменено
}

// String возвращает строковое представление счетчика
func (c Counter) String() string {
	return fmt.Sprintf("%s: %d", c.name, c.value)
}

func main() {
	// Создаем новый счетчик с именем "Visitors"
	visitors := NewCounter("Visitors")
	
	// Увеличиваем счетчик несколько раз
	visitors.Increment()
	visitors.Increment()
	visitors.Increment()
	
	// Выводим текущее значение счетчика
	fmt.Println(visitors)
	fmt.Println("Current value:", visitors.Value())
	
	// Уменьшаем счетчик
	visitors.Decrement()
	fmt.Println(visitors)
	
	// Сбрасываем счетчик
	visitors.Reset()
	fmt.Println(visitors)
	
	// Проверяем, что счетчик не опускается ниже нуля
	visitors.Decrement()
	fmt.Println(visitors)
}
