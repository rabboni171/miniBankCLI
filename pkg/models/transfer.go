package models

type Transfer struct {
	Sender    *Client
	Recipient *Client
	Amount    float64
}
