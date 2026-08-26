package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Чтение входных данных
    var maxRetriesStr string
    var successAttemptStr string
    fmt.Scanln(&maxRetriesStr)
    fmt.Scanln(&successAttemptStr)
    
    // Преобразование входных данных в целые числа
    maxRetries, _ := strconv.Atoi(maxRetriesStr)
    successAttempt, _ := strconv.Atoi(successAttemptStr)
    
    // Инициализация счетчика попыток
    attempt := 1
    
    // TODO: Напишите свой код ниже
    // Используйте оператор goto с меткой для реализации логики повторных попыток
    cycle:
        if attempt == successAttempt {
            fmt.Printf("Attempt %d succeeded\n", attempt)
            return
        }
        fmt.Printf("Attempt %d failed\n", attempt)
        attempt++
        if attempt > maxRetries {
            fmt.Println("All attempts failed")
            return
        }
        goto cycle
}
