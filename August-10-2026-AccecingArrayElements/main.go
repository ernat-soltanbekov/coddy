package main

import "fmt"

func main() {
    // Это наш массив фруктов
    fruits := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
    
    // TODO: Создайте переменную с именем firstFruit и присвойте ей первый элемент массива fruits
    firstFruit := fruits[0]
    // TODO: Создайте переменную с именем thirdFruit и присвойте ей третий элемент массива fruits
    thirdFruit := fruits[2]
    // Выведите фрукты
    fmt.Println("The first fruit is:", firstFruit)
    fmt.Println("The third fruit is:", thirdFruit)
}
