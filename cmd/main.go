package main

import (
	"fmt"
	"miniBankCLI/pkg"
	"time"
)

func main() {
	for {
		var choice int
		fmt.Println("1. Create a client")
		fmt.Println("2. Top up client's balance")
		fmt.Println("3. Show all clients")
		fmt.Println("4. Withdraw money from client's balance")
		fmt.Println("5. Stop session")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			pkg.CreateClient()
		case 2:
			pkg.TopUpClientBalance()
		case 3:
			pkg.ShowAllClients()
		case 4:
			pkg.TakeOffMoney()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		time.Sleep(2 * time.Second)
	}
}