package pkg

import "fmt"

// TakeOffMoney withdraw money from client's balance
func TakeOffMoney() {
	var name string
	var amount float64
	fmt.Print("Enter client name and amount to withdraw: ")
	fmt.Scan(&name, &amount)
	if balance, ok := Database[name]; ok {
		if balance >= amount {
			Database[name] = balance - amount
			fmt.Println("Well done!")
		} else {
			fmt.Println("Sorry, not enough money")
		}
	} else {
		fmt.Println("Client not found!")
	}
}
