package main

import "fmt"

func main() {
    // TODO: Инициализируйте массив с именем 'favoriteNumbers' этими 5 целыми числами: 7, 42, 8, 13, 99
    var favoriteNumbers = [...]int{7, 42, 8, 13, 99}
    // Вы можете использовать любой из вариантов синтаксиса:
    // favoriteNumbers := [5]int{7, 42, 8, 13, 99}
    // ИЛИ
    // var favoriteNumbers [5]int = [5]int{7, 42, 8, 13, 99}
    
    // Это выведет ваш массив
    fmt.Printf("My favorite numbers are: %v\n", favoriteNumbers)
}
