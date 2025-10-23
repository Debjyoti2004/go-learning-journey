package main

import (
	"fmt"
	"time"
)

type RestaurantBill struct {
	Name        string
	BillAmount  float32
	phoneNumber string
	CreatedAt   time.Time
	paymentMode string
}

func main() {
	RestaurantBill1 := RestaurantBill{
		Name:        "John Doe",
		BillAmount:  250.75,
		phoneNumber: "123-456-7890",
		CreatedAt:   time.Now(),
		paymentMode: "Credit Card",
	}

	fmt.Printf("Restaurant Bill Details:\n")
	fmt.Printf("Name: %s\n", RestaurantBill1.Name)
	fmt.Printf("Bill Amount: $%.2f\n", RestaurantBill1.BillAmount)
	fmt.Printf("Phone Number: %s\n", RestaurantBill1.phoneNumber)
	fmt.Printf("Created At: %s\n", RestaurantBill1.CreatedAt.Format(time.RFC1123))
	fmt.Printf("Payment Mode: %s\n", RestaurantBill1.paymentMode)
}
