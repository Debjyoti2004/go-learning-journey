package main

import "fmt"

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func ChangeOrderStatus(status OrderStatus) {
	switch status {
	case Received:
		fmt.Println("📦 Order Status: Received – Your order has been received.")
	case Confirmed:
		fmt.Println("✅ Order Status: Confirmed – Your order is confirmed.")
	case Prepared:
		fmt.Println("👨‍🍳 Order Status: Prepared – Your order is being prepared.")
	case Delivered:
		fmt.Println("🚚 Order Status: Delivered – Your order is out for delivery.")
	default:
		fmt.Println("❌ Unknown order status")
	}
}

func main() {
	ChangeOrderStatus(Received)
	ChangeOrderStatus(Confirmed)
	ChangeOrderStatus(Prepared)
	ChangeOrderStatus(Delivered)
}
