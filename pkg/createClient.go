package pkg

import "fmt"

//var database =

// CreateClient create a client
func CreateClient() {
	var name string
	var balance float64
	fmt.Print("Enter client name: ")
	fmt.Scan(&name)
	Database[name] = balance
	fmt.Println("_________")
	fmt.Println("Well done!")
	fmt.Println("_________")
}
