package pkg

import "fmt"

// TopUpClientBalance top up client's balance
func TopUpClientBalance() {
	var name string
	var amount float64
	fmt.Print("Enter client name and amount to top up: ")
	fmt.Scan(&name, &amount)
	if balance, ok := Database[name]; ok {
		Database[name] = balance + amount
		fmt.Println("Well done!")
	} else {
		fmt.Println("Client not found!")
	}
}
