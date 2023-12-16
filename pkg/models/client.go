package models

type Client struct {
	Name        string
	PhoneNumber string
	Balance     float64
	City        *City
}
