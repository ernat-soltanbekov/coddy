package main

import (
    "fmt"
    "strconv"
)

// TODO: Определите здесь вашу структуру Sensor
type Sensor struct {
    ID string
    Temperature float64
}

// TODO: Определите здесь ваш метод displayReading с получателем значения (value receiver)
func (t Sensor) displayReading() {
    fmt.Printf("Sensor %s: %.1f°C\n", t.ID, t.Temperature)
}
// TODO: Определите здесь ваш метод adjustTemperature с получателем значения (value receiver)
func (t Sensor) adjustTemperature (adjustment float64) {
    t.Temperature += adjustment
    fmt.Printf("Adjusted reading: %.1f°C\n", t.Temperature)
}

func main() {
    // Чтение входных данных
    var sensorID string
    var tempStr string
    var adjustStr string
    
    fmt.Scanln(&sensorID)
    fmt.Scanln(&tempStr)
    fmt.Scanln(&adjustStr)
    
    // Парсинг значений температуры и корректировки
    temperature, _ := strconv.ParseFloat(tempStr, 64)
    adjustment, _ := strconv.ParseFloat(adjustStr, 64)
    
    // TODO: Создайте экземпляр Sensor и вызовите необходимые методы
    sensor := Sensor {
        ID: sensorID,
        Temperature: temperature,
    }
    sensor.displayReading()
    sensor.adjustTemperature(adjustment)
    sensor.displayReading()
}
