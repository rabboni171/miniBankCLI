package pkg

import "fmt"

func TakeOffMoney() {
	var name string
	var amount float64

	fmt.Println("Client name: ")
	fmt.Scan(&name)
	var balance float64
	var has bool

	// check // has or no client
	for _, val := range Clients {
		if name == val.Name {
			balance = val.Balance
			has = true
		}
	}

	if !has {
		fmt.Println("Error! there is no such client in DB")
		return
	}

	fmt.Println("Type the sum to take off from balance: ")
	fmt.Scan(&amount)

	if balance < amount {
		fmt.Println("Error! Not enough money to take off!")
		return
	}
	for _, val := range Clients {
		if name == val.Name {
			val.Balance = balance - amount
		}
	}

}
