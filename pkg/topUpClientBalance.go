package pkg

import (
	"fmt"
)

// TopUpClientBalance top up client's balance
func TopUpClientBalance() {
	var name string
	var amount float64
	fmt.Print("Enter client name: ")
	fmt.Scan(&name)

	var balance float64
	var has bool

	// check// has or no client
	for _, val := range Clients {
		if name == val.Name {
			balance = val.Balance
			has = true
		}
	}

	if !has {
		fmt.Println("Error! there is no such client in DB")
	}

	fmt.Println("Sum to top up: ")
	fmt.Scan(&amount)

	for _, val := range Clients {
		if name == val.Name {
			val.Balance = balance + amount
		}
	}
}
