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
		fmt.Println("5. Transfer money ")
		fmt.Println("6. Profit of Bank")
		fmt.Println("7. Stop session")

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
			pkg.TransferMoney()
		case 6:
			pkg.ShowProfit()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		time.Sleep(1 * time.Second)
	}
}
