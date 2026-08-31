package main

import "fmt"

func main() {
    // Прочитать ввод
    var action string
    fmt.Scanln(&action)
    
    // TODO: Define your MediaPlayer interface here
    type MediaPlayer interface {
        Play() string
        Pause() string
        Stop() string
    }
    // TODO: Напишите ваш код ниже для обработки действия и вывода требуемого результата
    // Вывести результат
    // Remember to print in format: "MediaPlayer interface requires: [MethodName]() string"
    switch action {
    case "play":
    fmt.Printf("MediaPlayer interface requires: Play() string")
    case "pause":
    fmt.Printf("MediaPlayer interface requires: Pause() string")
    case "stop":
    fmt.Printf("MediaPlayer interface requires: Stop() string")
    }
}

