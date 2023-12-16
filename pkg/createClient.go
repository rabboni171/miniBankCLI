package pkg

import (
	"fmt"
	"miniBankCLI/pkg/models"
)

// CreateClient create a client
func CreateClient() {
	var name, cityName string
	var phoneNumber string

	fmt.Print("Enter client name: ")
	fmt.Scan(&name)

	fmt.Print("Enter client phone number: ")
	fmt.Scan(&phoneNumber)

	fmt.Println("Enter client city name")
	fmt.Scan(&cityName)

	var city models.City
	var has bool

	for _, v := range Cities {
		if v.Name == cityName {
			city = v
			has = true
		}
	}

	if !has {
		fmt.Println("The city has not found in DB, call 544")
		return
	}

	Clients = append(Clients, &models.Client{
		Name:        name,
		PhoneNumber: phoneNumber,
		Balance:     0.0,
		City:        &city,
	})

	fmt.Println("__________")
	fmt.Println("Well done")
	fmt.Println("__________")

}
