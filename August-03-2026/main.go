package main

import "fmt"

// Эта функция выводит приветственное сообщение
func sayHello() {
    fmt.Println("Hello, friend!")
}

// Эта функция выводит прощальное сообщение
func sayGoodbye() {
    fmt.Println("Goodbye, friend!")
}

// Эта функция выводит сообщение с благодарностью
func sayThankYou() {
    fmt.Println("Thank you, friend!")
}

func main() {
    // TODO: Вызовите функцию sayHello
    sayHello()
    // TODO: Вызовите функцию sayThankYou
    sayThankYou()
    // TODO: Вызовите функцию sayGoodbye
    sayGoodbye()
}
