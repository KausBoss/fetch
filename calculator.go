package main

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

func calculatePoints(receipt *Receipt) (points int) {
	points = 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	points += retailNamePoints(receipt.Retailer)

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	points += totalFloatPoints(receipt.Total)

	// Rule 3: 25 points if the total is a multiple of 0.25.
	points += totalDivisiblePoints(receipt.Total)

	// Rule 4: 5 points for every two items on the receipt.
	points += itemCountPoints(receipt.Items)

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	points += itemDescriptionPoints(receipt.Items)

	// Rule 6: 6 points if the day in the purchase date is odd.
	points += purchaseDatePoints(receipt.PurchaseDate)

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	points += purchaseTimePoints(receipt.PurchaseTime)

	return points
}

// retailNamePoints One point for every alphanumeric character in the retailer name.
func retailNamePoints(retailName string) (points int) {
	points = 0
	for ch := range retailName {
		if unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch)) {
			points++
		}
	}
	return
}

// totalFloatPoints 50 points if the total is a round dollar amount with no cents.
func totalFloatPoints(total string) (points int) {
	points = 0
	totalFloat, err := strconv.ParseFloat(total, 64)
	if err == nil && totalFloat == float64(int(totalFloat)) {
		points = 50
	}
	return
}

// totalDivisiblePoints 25 points if the total is a multiple of 0.25.
func totalDivisiblePoints(total string) (points int) {
	points = 0
	totalCents, err := strconv.ParseFloat(total, 64)
	if err == nil && int(totalCents*100)%25 == 0 {
		points = 25
	}
	return
}

// itemCountPoints 5 points for every two items on the receipt.
func itemCountPoints(items []Item) int {
	return (len(items) / 2) * 5
}

// itemDescriptionPoints If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
func itemDescriptionPoints(items []Item) int {
	points := 0
	for _, item := range items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				// rounding up to nearest integer
				points += int((priceFloat * 0.2) + 0.5)
			}
		}
	}
	return points
}

// purchaseDatePoints 6 points if the day in the purchase date is odd.
func purchaseDatePoints(date string) (points int) {
	points = 0
	purchaseDate, err := time.Parse("2006-01-02", date)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}
	return points
}

// purchaseTimePoints 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func purchaseTimePoints(timeStr string) (points int) {
	points = 0
	purchaseTime, err := time.Parse("15:04", timeStr)

	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points = 10
	}
	if err == nil && purchaseTime.Hour() == 16 && purchaseTime.Second() == 00 {
		points = 10
	}
	return
}
