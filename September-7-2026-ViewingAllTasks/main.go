package main

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
    )

type Task struct {
    Name string
    Completed bool
}

func viewAllTasks(parametr []Task) {
    for i := 0; i < len(parametr); i++ {
        if parametr[i].Completed == true {
            fmt.Printf("[x] %s\n", parametr[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", parametr[i].Name)
        }
    }
}

func main() {
    var divideTask []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scannedText, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Invalid string")
        return
    }
    scanner.Scan()
    secondString := scanner.Text()
    dividedString := strings.Split(secondString, ",")
    for i := 0; i < scannedText; i++ {
        currentNote := dividedString[i]
        dividedNote := strings.Split(currentNote, ":")
        boolDefine, err := strconv.ParseBool(dividedNote[1])
        if err != nil {
            fmt.Println("Shit happens. Try again later.")
            return
        }
        divideTask = append(divideTask, Task{dividedNote[0], boolDefine})
    }
    completedCount := 0
    for count := 0; count < len(divideTask); count++ {
        if divideTask[count].Completed == true {
            completedCount++
        }
    }
    viewAllTasks(divideTask)
    fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", scannedText, completedCount, scannedText - completedCount)
}
