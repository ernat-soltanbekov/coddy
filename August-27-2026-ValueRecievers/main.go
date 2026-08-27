package main // Объявляем пакет main, потому что выполнение программы начинается с функции main

import ( // Подключаем пакеты, которые понадобятся ниже
    "fmt" // Используем для чтения данных из терминала и вывода результатов
    "strconv" // Используем для преобразования строк в числа
)

type Sensor struct { // Создаём структуру Sensor, чтобы хранить ID и температуру одного датчика
    ID string // Поле ID хранит идентификатор датчика в виде строки
    Temperature float64 // Поле Temperature хранит температуру как число с плавающей точкой
}

func (r Sensor) displayReading() { // Создаём метод с value receiver, поэтому r является копией Sensor
    fmt.Printf("Sensor %s: %.1f°C\n", r.ID, r.Temperature) // Выводим ID и температуру; %.1f оставляет один знак после точки
}

func (t Sensor) adjustTemperature(adjustment float64) { // Создаём метод с value receiver и принимаем величину корректировки
    t.Temperature += adjustment // Прибавляем корректировку к температуре копии, поэтому оригинальный Sensor не изменяется
    fmt.Printf("Adjusted reading: %.1f°C\n", t.Temperature) // Выводим изменённую температуру из копии с одним знаком после точки
}

func main() { // Главная функция, с которой начинается выполнение программы
    var sensorID string // Создаём переменную для хранения ID датчика, который введёт пользователь
    var tempStr string // Создаём строку для температуры, потому что из терминала она сначала приходит как текст
    var adjustStr string // Создаём строку для корректировки, потому что она тоже сначала приходит как текст
    
    fmt.Scanln(&sensorID) // Считываем ID датчика из терминала и записываем его в sensorID
    fmt.Scanln(&tempStr) // Считываем температуру из терминала как строку
    fmt.Scanln(&adjustStr) // Считываем корректировку из терминала как строку
    
    temperature, _ := strconv.ParseFloat(tempStr, 64) // Преобразуем строку температуры в число float64; ошибку здесь игнорируем
    adjustment, _ := strconv.ParseFloat(adjustStr, 64) // Преобразуем строку корректировки в число float64
    
    sensor := Sensor{ // Создаём экземпляр структуры Sensor с полученными данными
        ID: sensorID, // Передаём введённый ID в поле ID
        Temperature: temperature, // Передаём преобразованную температуру в поле Temperature
    }
    
    sensor.displayReading() // Показываем исходную температуру датчика
    sensor.adjustTemperature(adjustment) // Передаём корректировку методу, который изменяет только копию sensor
    sensor.displayReading() // Снова показываем оригинальный sensor, температура которого осталась неизменной
}
