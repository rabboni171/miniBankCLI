package pkg

import "fmt"

// ShowAllClients show all clients
func ShowAllClients() {

	var name string
	fmt.Println("Name of client: ")

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

	fmt.Printf("client's balance %v is %v \n", name, balance)
}
