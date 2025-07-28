package main

import "fmt"

type paymenter interface {
	pay(amount float32, name string)
	refund(amount float32, account string, name string)
}

type payment struct {
	Name    string
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount, p.Name)
}

func (p payment) refundPayment(amount float32, account string) {
	p.gateway.refund(amount, account, p.Name)
}

type paypal struct{}

func (p paypal) pay(amount float32, name string) {
	fmt.Println("====================================")
	fmt.Println("|        Payment via PayPal        |")
	fmt.Println("====================================")
	fmt.Printf("| %-15s | %-15s |\n", "User", name)
	fmt.Printf("| %-15s | %-15.2f |\n", "Amount", amount)
	fmt.Println("====================================")
}

func (p paypal) refund(amount float32, account string, name string) {
	fmt.Println("====================================")
	fmt.Println("|        Refund via PayPal         |")
	fmt.Println("====================================")
	fmt.Printf("| %-15s | %-15s |\n", "User", name)
	fmt.Printf("| %-15s | %-15s |\n", "Account", account)
	fmt.Printf("| %-15s | %-15.2f |\n", "Amount", amount)
	fmt.Println("====================================")
}

type phonepay struct{}

func (p phonepay) pay(amount float32, name string) {
	fmt.Println("====================================")
	fmt.Println("|       Payment via PhonePe        |")
	fmt.Println("====================================")
	fmt.Printf("| %-15s | %-15s |\n", "User", name)
	fmt.Printf("| %-15s | %-15.2f |\n", "Amount", amount)
	fmt.Println("====================================")
}

func (p phonepay) refund(amount float32, account string, name string) {
	fmt.Println("====================================")
	fmt.Println("|       Refund via PhonePe         |")
	fmt.Println("====================================")
	fmt.Printf("| %-15s | %-15s |\n", "User", name)
	fmt.Printf("| %-15s | %-15s |\n", "Account", account)
	fmt.Printf("| %-15s | %-15.2f |\n", "Amount", amount)
	fmt.Println("====================================")
}

func main() {
	var name string
	fmt.Print("Enter Your name: ")
	fmt.Scan(&name)

	fmt.Println("\n--- PayPal Transactions ---")
	paypalGW := paypal{}
	user1 := payment{
		gateway: paypalGW,
		Name:    name,
	}
	user1.makePayment(123.44)
	user1.refundPayment(100.01, "debjyoti@gmail.com")

	fmt.Println("\n--- PhonePe Transactions ---")
	phonepayGW := phonepay{}
	user2 := payment{
		gateway: phonepayGW,
		Name:    name,
	}
	user2.makePayment(555.55)
	user2.refundPayment(200.50, "debjyoti@phonepe")
}
