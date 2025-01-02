package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

const (
	DepositOption            = 1
	WithdrawOption           = 2
	ViewBalanceOption        = 3
	TransactionHistoryOption = 4
	ExitOption               = 5
)

func FindAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("Account not found")
}
func Deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be greater than 0")
	}
	account.Balance += amount
	transaction := fmt.Sprintf("Deposited: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}
func Withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be greater than 0")
	}
	if amount > account.Balance {
		return errors.New("Insufficient balance")
	}
	account.Balance -= amount
	transaction := fmt.Sprintf("Withdrew: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}

func DisplayTransactionHistory(account *Account) {
	fmt.Println("Transaction History:")
	for _, transaction := range account.TransactionHistory {
		fmt.Println(transaction)
	}
}

func main() {
	accounts = append(accounts, Account{ID: 1, Name: "abc", Balance: 5200})
	accounts = append(accounts, Account{ID: 2, Name: "def", Balance: 8990})
	accounts = append(accounts, Account{ID: 3, Name: "ghi", Balance: 4224})
	accounts = append(accounts, Account{ID: 4, Name: "jhi", Balance: 6773})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. View Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		if choice == ExitOption {
			break
		}

		var accountID int
		fmt.Print("Enter Account ID: ")
		fmt.Scan(&accountID)

		account, err := FindAccountByID(accountID)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch choice {
		case DepositOption:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := Deposit(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Successfully deposited $%.2f. New Balance: %.2f\n", amount, account.Balance)
			}

		case WithdrawOption:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := Withdraw(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Successfully withdrew $%.2f. New Balance: %.2f\n", amount, account.Balance)
			}

		case ViewBalanceOption:
			fmt.Printf("Account Balance: %.2f\n", account.Balance)

		case TransactionHistoryOption:
			DisplayTransactionHistory(account)

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
