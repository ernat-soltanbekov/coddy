package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: Напишите свой код ниже
// 1. Определите интерфейс Speaker
type Speaker interface {
    Speak() string
}
// 2. Определите структуры Person и Parrot
type Person struct {
    Name string
}
type Parrot struct {
    Name string
}
// 3. Реализуйте методы Speak() для обеих структур
func (a Person) Speak() string {
    return fmt.Sprintf("Hello, I'm %s", a.Name)
}
func (a Parrot) Speak() string {
    return fmt.Sprintf("Squawk! %s says hello!", a.Name)
}
// 4. Создайте функцию makeAllSpeak
func makeAllSpeak(kusok []Speaker) {
    for _, element := range kusok {
        fmt.Println(element.Speak())
    }
}

func main() {
	// Чтение входных данных
	var numSpeakersStr string
	var speakerTypesStr string
	var speakerNamesStr string
	
	fmt.Scanln(&numSpeakersStr)
	fmt.Scanln(&speakerTypesStr)
	fmt.Scanln(&speakerNamesStr)
	
	// Парсинг входных данных
	numSpeakers, _ := strconv.Atoi(numSpeakersStr)
	speakerTypes := strings.Split(speakerTypesStr, ",")
	speakerNames := strings.Split(speakerNamesStr, ",")	
    // 5. Создайте объекты speaker на основе входных данных и сохраните их в срез
    speakers := []Speaker{}
    for i := 0; i < numSpeakers; i++ {
        switch speakerTypes[i] {
            case "person":
            speakers = append(speakers, Person{Name: speakerNames[i]})
            case "parrot":
            speakers = append(speakers, Parrot{Name: speakerNames[i]})
        }
    }
    // 6. Вызовите makeAllSpeak с вашим срезом
    makeAllSpeak(speakers)
}
