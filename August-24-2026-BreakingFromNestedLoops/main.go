package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // Чтение входных данных
    var dimensions string
    var targetStr string
    fmt.Scanln(&dimensions)
    fmt.Scanln(&targetStr)
    
    // Разбор размеров
    dimParts := strings.Split(dimensions, ",")
    rows, _ := strconv.Atoi(dimParts[0])
    cols, _ := strconv.Atoi(dimParts[1])
    
    // Разбор целевого числа
    target, _ := strconv.Atoi(targetStr)
    
    // Предопределенная сетка
    grid := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    
    // TODO: Напишите свой код ниже
    // Используйте вложенные циклы с именованным break для поиска цели
    found := false
    search:
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if grid[row][col] == target {
            fmt.Printf("Found %d at position (%d, %d)\n", target, row, col)
            found = true
            break search
            }
        }
    }
    if !found {
    fmt.Printf("Target %d not found", target)
    }
}
