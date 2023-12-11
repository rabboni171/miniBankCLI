package pkg

import "fmt"

func TransferMoney() {
	var sender, recipient string
	var amount float64

	fmt.Print("Sender's name: ")
	fmt.Scan(&sender)
	balanceSender, ok := Database[sender]

	if !ok {
		fmt.Println("There is not balance of sender")
	}
	fmt.Print("Recipient's name")
	fmt.Scan(&recipient)
	balanceRecipient, ok := Database[recipient]

	if !ok {
		fmt.Println("There is not balance of recipient")
	}
	fmt.Println("What is the amount to transfer?: ")
	fmt.Scan(&amount)
	if balanceSender < amount {
		fmt.Println("Not enough money")
		return
	}
	var comission float64 = amount / 100 * Percent
	Database[sender] = balanceSender - amount
	Database[recipient] = (balanceRecipient + amount) - comission

	balanceProfit, ok := Database["Profit"]
	if !ok {
		Database["Profit"] = 0
	}

	Database["Profit"] = balanceProfit + comission

	fmt.Println("Succes")
}
