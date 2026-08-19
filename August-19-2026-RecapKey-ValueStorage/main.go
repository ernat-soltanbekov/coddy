package main

import (
	"fmt"
	"sort"
)

// addWord добавляет слово в карту счетчика
func addWord(counter map[string]int, word string) {
	// Ваш код здесь
    counter[word]++
}

// getCount возвращает количество для конкретного слова
func getCount(counter map[string]int, word string) int {
	// Ваш код здесь
	return counter[word]
}

// removeWord удаляет слово из счетчика
// Возвращает true, если слово было найдено и удалено
func removeWord(counter map[string]int, word string) bool {
	// Ваш код здесь
    _, exists := counter[word]
    if exists {
        delete(counter, word)
        return true
    }
    return false
}

// printWordCounts выводит количество слов в алфавитном порядке
func printWordCounts(counter map[string]int) {
	// Получить все ключи (слова) из карты
	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	
	// Отсортировать слова по алфавиту
	sort.Strings(words)
	
	// Вывести каждое слово и его количество в отсортированном порядке
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, counter[word])
	}
}

func main() {
	// Инициализировать карту счетчика слов
	wordCounter := make(map[string]int)
	
	// Добавить несколько слов
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	for _, word := range words {
		addWord(wordCounter, word)
	}
	
	// Вывести количество слов
	fmt.Println("Word counts:")
	printWordCounts(wordCounter)
	
	// Получить количество для конкретных слов
	fmt.Printf("\nCount for 'apple': %d\n", getCount(wordCounter, "apple"))
	fmt.Printf("Count for 'grape': %d\n", getCount(wordCounter, "grape"))
	
	// Удалить слово
	removed := removeWord(wordCounter, "banana")
	fmt.Printf("\nRemoved 'banana': %v\n", removed)
	
	// Попробовать удалить несуществующее слово
	removed = removeWord(wordCounter, "grape")
	fmt.Printf("Removed 'grape': %v\n", removed)
	
	// Вывести итоговое количество слов
	fmt.Println("\nFinal word counts:")
	printWordCounts(wordCounter)
}
