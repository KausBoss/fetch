package main

import "github.com/google/uuid"

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        string    `json:"total"`
	Points       int       `json:"points,omitempty"`
	ID           uuid.UUID `json:"id,omitempty"`
}

func (r Receipt) SetPoints() {
	r.Points = calculatePoints(&r)
}
