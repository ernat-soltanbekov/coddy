package main

import (
	"fmt"
)

// Эта функция имитирует открытие файла
// Она возвращает ошибку, если файл не существует
func openFile(filename string) error {
	if filename == "data.txt" {
		return nil // nil означает отсутствие ошибки
	} else {
		return fmt.Errorf("file %s not found", filename)
	}
}

func main() {
	filename := "config.txt"
	
	// Вызов функции openFile
	err := openFile(filename)
	
	// TODO: Проверьте, не является ли err отличным от nil (что означает наличие ошибки)
	// Если возникла ошибка, выведите: "Error: " с последующим сообщением об ошибке
	// Если ошибки не было, выведите: "File opened successfully"
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File opened successfully")
}
