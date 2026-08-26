package main // Главный пакет программы, с него начинается выполнение

import (
	"fmt" // Нужен для чтения данных и вывода результата
	"strconv" // Нужен для преобразования строк в числа
)

type Account struct { // Создаём структуру счёта, чтобы хранить его данные вместе
	Name string // Имя владельца счёта хранится как строка
	Balance float64 // Баланс хранится как число с плавающей точкой
}

func (d *Account) deposit(amount float64) { // Pointer receiver нужен, чтобы изменить настоящий баланс; amount — сумма пополнения
	d.Balance += amount // Прибавляем сумму пополнения непосредственно к балансу оригинального счёта
}

func (w *Account) withdraw(amount float64) { // Pointer receiver нужен, чтобы изменить настоящий баланс; amount — сумма снятия
	w.Balance -= amount // Вычитаем сумму снятия из баланса оригинального счёта
}

func (b Account) displayBalance() { // Value receiver получает копию счёта, потому что этот метод ничего не изменяет
	fmt.Printf("Account: %s, Balance: $%.2f\n", b.Name, b.Balance) // Выводим имя и баланс, показывая ровно два знака после точки
}

func main() { // Главная функция, с которой начинается выполнение программы
	var name string // Переменная для имени владельца счёта
	var initialBalanceStr string // Строка для первоначального баланса, который пока ещё является текстом
	var transactionAmountStr string // Строка для суммы транзакции, которая пока ещё является текстом

	fmt.Scanln(&name) // Читаем имя владельца из терминала
	fmt.Scanln(&initialBalanceStr) // Читаем первоначальный баланс как строку
	fmt.Scanln(&transactionAmountStr) // Читаем сумму транзакции как строку

	initialBalance, _ := strconv.ParseFloat(initialBalanceStr, 64) // Преобразуем строку начального баланса в число float64
	transactionAmount, _ := strconv.ParseFloat(transactionAmountStr, 64) // Преобразуем строку транзакции в число float64

	account := Account{ // Создаём конкретный экземпляр структуры Account
		Name: name, // Записываем введённое имя в поле Name
		Balance: initialBalance, // Записываем начальный баланс в поле Balance
	}

	account.displayBalance() // Показываем первоначальный баланс
	account.deposit(transactionAmount) // Передаём сумму в deposit, который увеличивает настоящий баланс
	account.displayBalance() // Показываем баланс после пополнения
	account.withdraw(transactionAmount) // Передаём сумму в withdraw, который уменьшает настоящий баланс
	account.displayBalance() // Показываем итоговый баланс
}
