package main

import "fmt"

func main() {
    // Чтение входных данных
    var alertLevel string
    fmt.Scanln(&alertLevel)
    
    // TODO: Напишите свой код ниже
    // Создайте оператор `switch` с использованием `fallthrough` для системы оповещения
    switch alertLevel {
    case "critical":
        fmt.Println("CRITICAL: System shutdown imminent!")
        fallthrough
    case "high":
        fmt.Println("HIGH: Immediate attention required!")
        fallthrough
    case "medium":
        fmt.Println("MEDIUM: Please review when possible")
    case "low":
        fmt.Println("LOW: Informational only")
    }
}
