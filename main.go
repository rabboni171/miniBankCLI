package main

import (
	"fmt"
	"time"
)

var database = make(map[string]int)

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
			createClient()
		case 2:
			topUpClientBalance()
		case 3:
			showAllClients()
		case 4:
			withdrawMoneyFromBalance()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		time.Sleep(2 * time.Second)
	}
}

// create a client
func createClient() {
	var name string
	var balance int
	fmt.Print("Enter client name and initial balance: ")
	fmt.Scan(&name, &balance)
	database[name] = balance
	fmt.Println("_________")
	fmt.Println("Well done!")
	fmt.Println("_________")

}

// top up client's balance
func topUpClientBalance() {
	var name string
	var sum int
	fmt.Print("Enter client name and amount to top up: ")
	fmt.Scan(&name, &sum)
	if balance, ok := database[name]; ok {
		database[name] = balance + sum
		fmt.Println("Well done!")
	} else {
		fmt.Println("Client not found!")
	}
}

// show all clients
func showAllClients() {
	fmt.Println("Clients:")
	for k, v := range database {
		if len(database) == 0 {
			fmt.Println("No clients found.")
			return
		}
		fmt.Println(k, v)
	}
}

// withdraw money from client's balance
func withdrawMoneyFromBalance() {
	var name string
	var someMoney int
	fmt.Print("Enter client name and amount to withdraw: ")
	fmt.Scan(&name, &someMoney)
	if balance, ok := database[name]; ok {
		if balance >= someMoney {
			database[name] = balance - someMoney
			fmt.Println("Well done!")
		} else {
			fmt.Println("Sorry, not enough money")
		}
	} else {
		fmt.Println("Client not found!")
	}
}
