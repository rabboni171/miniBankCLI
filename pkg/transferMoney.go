package pkg

import (
	"fmt"
	"miniBankCLI/pkg/models"
)

// Проверка двух счетов
// Подсчет комиссии
// Перевод (
//	Снятие денег с 1ого счета,
//	Пополнение денег на 2ой счет,
//	Пополенение Прибыль в банк
//)

func TransferMoney() {
	var sender, recipient string
	var amount float64

	fmt.Print("Sender's name: ")
	fmt.Scan(&sender)

	var Sender *models.Client
	var has bool

	// check // has or no sender
	for _, val := range Clients {
		if sender == val.Name {
			has = true
			Sender = val
		}

	}

	if !has {
		fmt.Println("No senders's account")
		return
	}

	fmt.Println("Recepient's name: ")
	fmt.Scan(&recipient)

	has = false
	var Recipient *models.Client

	//  проверка на наличие клиента
	for _, val := range Clients {
		if recipient == val.Name {
			has = true
			Recipient = val
		}
	}

	if !has {
		fmt.Println("No sender's account")
		return
	}

	fmt.Println("sum for transfer:")
	fmt.Scan(&amount)

	if Sender.Balance < amount {
		fmt.Println("not enough money")
		return
	}
	var comission float64 = amount / 100 * Percent

	for _, val := range Clients {
		if sender == val.Name {
			val.Balance = Sender.Balance
		}
	}

	for _, val := range Clients {
		if recipient == val.Name {
			val.Balance = Recipient.Balance + amount - comission
		}
	}
	has = false
	var balanceProfit float64
	for _, val := range Clients {
		if "profit" == val.Name {
			balanceProfit = val.Balance
			has = true
		}
	}

	if !has {
		Clients = append(Clients, &models.Client{
			Name:        "profit",
			PhoneNumber: "554",
			Balance:     0,
		})
	}

	for _, val := range Clients {
		if "profit" == val.Name {
			val.Balance = balanceProfit + comission
		}
	}

	Transfers = append(Transfers, models.Transfer{
		Sender:    Sender,
		Recipient: Recipient,
		Amount:    amount,
	})

	fmt.Println("Success")
}
