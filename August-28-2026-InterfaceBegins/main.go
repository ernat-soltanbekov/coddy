package main                                      // Объявляем пакет main, потому что это исполняемая программа.

import "fmt"                                      // Подключаем fmt, чтобы читать ввод и выводить результат в терминал.

func main() {                                     // Запускаем главную функцию программы.
    var action string                             // Создаём переменную для хранения действия: play, pause или stop.
    fmt.Scanln(&action)                           // Читаем действие пользователя из терминала и записываем его в action.

    type MediaPlayer interface {                  // Определяем интерфейс MediaPlayer как контракт для медиаплеера.
        Play() string                              // Требуем от любого MediaPlayer метод Play, который возвращает string.
        Pause() string                             // Требуем от любого MediaPlayer метод Pause, который возвращает string.
        Stop() string                              // Требуем от любого MediaPlayer метод Stop, который возвращает string.
    }

    var methodName string                          // Создаём переменную для названия метода, который соответствует введённому действию.

    switch action {                                // Проверяем, какое действие пользователь ввёл.
    case "play":                                   // Если пользователь ввёл "play", выбираем метод Play.
        methodName = "Play"                        // Сохраняем название требуемого метода Play.
    case "pause":                                  // Если пользователь ввёл "pause", выбираем метод Pause.
        methodName = "Pause"                       // Сохраняем название требуемого метода Pause.
    case "stop":                                   // Если пользователь ввёл "stop", выбираем метод Stop.
        methodName = "Stop"                        // Сохраняем название требуемого метода Stop.
    }

    fmt.Printf("MediaPlayer interface requires: %s() string\n", methodName) // Выводим название метода и его требуемый тип возвращаемого значения.
}
