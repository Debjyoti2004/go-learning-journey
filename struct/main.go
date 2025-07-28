package main

import (
	"fmt"
	"time"
)

type order struct{
	id string
	customerName string
	amount float64
	orderAt time.Time
	status string
}

// constructor
func newOrder(id string, customerName string, amount float64) order{
	return order{
		id: id,
		customerName: customerName,
		amount: amount,
		status: "pending",
		orderAt: time.Now(),	
	}
}

// Reciver Type 
// when you update the value then use pointer
func (o order) getOrderDetails() string{
	return fmt.Sprintf("Order ID: %s, Customer: %s, Amount: $%.2f, Status: %s, Order At: %s", o.id, o.customerName, o.amount, o.status, o.orderAt)
}

func (o *order) changeStatus(newStatus string){
	o.status = newStatus
}


func main(){

	// order1 := order{
	// 	id: "123456789",
	// 	customerName: "Debjyoti",
	// 	amount: 100.00,
	// 	status: "pending",
	// }

	// order1.orderAt = time.Now()

	// order1.changeStatus("shipped")

	// fmt.Println(order1.getOrderDetails())


	order1 := newOrder("123456789", "Debjyoti", 100.00)
	fmt.Println(order1.getOrderDetails())
	order1.changeStatus("shipped")
	fmt.Println(order1.getOrderDetails())

}