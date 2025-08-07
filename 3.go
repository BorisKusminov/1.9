package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма для пополнения должна быть положительной")
		return
	}
	b.Balance += amount
	fmt.Printf("Счет пополнен на %.2f. Новый баланс: %.2f\n", amount, b.Balance)
}

// Метод для снятия средств
func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Сумма для снятия должна быть положительной")
		return
	}
	if b.Balance < amount {
		fmt.Printf("Недостаточно средств. Текущий баланс: %.2f, запрошенная сумма: %.2f\n", b.Balance, amount)
		return
	}
	b.Balance -= amount
	fmt.Printf("Со счета снято %.2f. Новый баланс: %.2f\n", amount, b.Balance)
}

func main() {
	// Создаем банковский счет
	account := BankAccount{
		Owner:   "Иван Петров",
		Balance: 1000.00,
	}

	fmt.Printf("Создан счет. Владелец: %s, начальный баланс: %.2f\n", account.Owner, account.Balance)

	// Демонстрация работы методов
	account.Deposit(500.50)
	account.Withdraw(200.75)
	account.Withdraw(1500.00) // Попытка снять больше, чем есть на счете
	account.Deposit(-100.00)  // Некорректная сумма для пополнения
	account.Withdraw(0.00)    // Некорректная сумма для снятия
}
