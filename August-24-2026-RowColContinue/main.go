package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Чтение входных данных
    var rowsInput string
    var skipConditionInput string
    fmt.Scanln(&rowsInput)
    fmt.Scanln(&skipConditionInput)
    
    // Разбор входных данных
    numRows, _ := strconv.Atoi(rowsInput)
    skipCondition, _ := strconv.Atoi(skipConditionInput)
    
    // Предоставленный двумерный массив данных
    data := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }
    
    // TODO: Напишите свой код ниже
    // Используйте вложенные циклы с именованным continue для обработки данных
    // Проверьте каждую строку на наличие числа условия пропуска
    // Выведите соответствующие сообщения в зависимости от того, найдено ли условие
    processRows:
for row := 0; row < numRows; row++ {
    for col := 0; col < 5; col++ {
        if data[row][col] == skipCondition {
            fmt.Printf("Skipping row %d due to condition", row)
            continue processRows
        }
    }

    fmt.Printf("Processing row %d: %d %d %d %d %d\n",
        row,
        data[row][0],
        data[row][1],
        data[row][2],
        data[row][3],
        data[row][4],
    )
}
}
