package pkg

import (
	"fmt"
	"miniBankCLI/pkg/models"
)

// ShowProfit displays the profit of the bank
func ShowProfit() {
	var balance float64
	var has bool
	// check for client
	for _, val := range Clients {
		if "profit" == val.Name {
			balance = val.Balance
			has = true
		}

	}

	if !has {
		Clients = append(Clients, &models.Client{
			Name:        "profit",
			Balance:     0.0,
			PhoneNumber: "544",
			City:        &Cities[0],
		})
	}

	fmt.Println(balance)
}
