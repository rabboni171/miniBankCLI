package pkg

import "fmt"

// show all clients
func ShowAllClients() {
	fmt.Println("Clients:")
	for k, v := range Database {
		if len(Database) == 0 {
			fmt.Println("No clients found.")
			return
		}
		fmt.Println(k, v)
	}
}
