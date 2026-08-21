package main
import (
    "errors"
    "fmt"
)
// divide делит a на b и возвращает результат
// Если b равно 0, функция возвращает ошибку
func divide(a, b int) (int, error) {
    // TODO: Реализовать функцию divide
    // Если b равно 0, вернуть 0 и ошибку с сообщением "division by zero"
    // В противном случае вернуть a/b и nil в качестве ошибки
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
func main() {
    // Первый тестовый случай
    a, b := 10, 2
    result, err := divide(a, b)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    } else {
        fmt.Printf("Result: %d\n", result)
    }

    // Второй тестовый случай
    a, b = 8, 0
    result, err = divide(a, b)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
    } else {
        fmt.Printf("Result: %d\n", result)
    }
}
