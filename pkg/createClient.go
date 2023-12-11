package pkg

import "fmt"

//var database =

// create a client
func CreateClient() {
	var name string
	var balance int
	fmt.Print("Enter client name: ")
	fmt.Scan(&name)
	Database[name] = balance
	fmt.Println("_________")
	fmt.Println("Well done!")
	fmt.Println("_________")
}
