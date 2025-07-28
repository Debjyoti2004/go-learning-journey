package main

import "fmt"

type paypal struct {
	Name string
}

func (p paypal) pay(amount float32) {
	fmt.Println("Wellcome To The Paypal")
	fmt.Println("-----------------------")
	fmt.Println("There User Name is: ", p.Name)
	fmt.Printf("The Amount is: %.2f\n ", amount)
}

type phonepay struct {
	Name string
}

func (p phonepay) pay(amount float32) {
	fmt.Println("Wellcome To The Phonepay")
	fmt.Println("-----------------------")
	fmt.Println("There User Name is: ", p.Name)
	fmt.Printf("The Amount is: %.2f\n ", amount)
}

func main() {
	var name string
	fmt.Println("Enter Your name ")
	fmt.Scan(&name)
	user := paypal{Name: name}

	user.pay(199.00)

}
