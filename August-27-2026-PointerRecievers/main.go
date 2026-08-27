package main

import (
	"fmt"                              // fmt нужен для вывода текста в терминал
	"strconv"                          // strconv нужен для преобразования string в float64
)

type Account struct {                 // Объявляем структуру банковского счёта
	Name string                         // Name хранит имя владельца счёта
	Balance float64                     // Balance хранит баланс счёта
}

func (a *Account) deposit(amount float64) { // a — указатель на счёт, amount — сумма пополнения
	a.Balance += amount                    // Увеличиваем настоящий Balance на значение amount
}

func (a *Account) withdraw(amount float64) { // a — указатель на счёт, amount — сумма снятия
	a.Balance -= amount                       // Уменьшаем настоящий Balance на значение amount
}

func (a Account) displayBalance() {      // Account без * означает получатель-значение, то есть копия
	fmt.Printf("Account: %s, Balance: $%.2f\n", a.Name, a.Balance) // Выводим имя и баланс с 2 знаками после точки
}

func main() {                             // Здесь начинается выполнение программы
	var name string                      // Переменная для имени владельца
	var initialBalanceStr string          // Здесь первоначальный баланс пока хранится как строка
	var transactionAmountStr string       // Здесь сумма операции пока хранится как строка
	
	fmt.Scanln(&name)                     // Читаем имя из терминала
	fmt.Scanln(&initialBalanceStr)        // Читаем первоначальный баланс как строку
	fmt.Scanln(&transactionAmountStr)     // Читаем сумму операции как строку
	
	initialBalance, _ := strconv.ParseFloat(initialBalanceStr, 64) // Превращаем строку баланса в float64
	transactionAmount, _ := strconv.ParseFloat(transactionAmountStr, 64) // Превращаем строку суммы в float64
	
	account := Account{                    // Создаём экземпляр структуры Account
		Name:    name,                      // Кладём имя из переменной name в поле Name
		Balance: initialBalance,            // Кладём начальный баланс в поле Balance
	}
	
	account.displayBalance()               // Показываем первоначальный баланс
	account.deposit(transactionAmount)     // Передаём transactionAmount в параметр amount метода deposit
	account.displayBalance()               // Показываем баланс после пополнения
	account.withdraw(transactionAmount)     // Передаём transactionAmount в параметр amount метода withdraw
	account.displayBalance()               // Показываем итоговый баланс
}
