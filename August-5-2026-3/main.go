package main

import "fmt"

// Глобальная переменная
var message string = "Hello from global scope"

func updateMessage() {
    // Это создает новую локальную переменную вместо изменения глобальной
    msg := "Hello from function scope"
    fmt.Println("Inside function:", msg)
    message = msg
}

func main() {
    fmt.Println("Before function call:", message)
    updateMessage()
    
    // TODO: Глобальное сообщение все еще имеет исходное значение
    // Измените функцию updateMessage(), чтобы изменить глобальную переменную
    fmt.Println("After function call:", message)
}
