// This is for call the function

package main
import "fmt"

func main (){
	fmt.Println("This function return this value:",sayHello())
	fmt.Println("The sum of this numbrt is this :",add(2,5))
	fmt.Println("The table of this number is this :",table(21,183,"Debjyoti"))

	// For multiple return value
	sum,sub,mul,div := calculater(10,5)
	fmt.Println("The sum of this number is this :",sum)
	fmt.Println("The sub of this number is this :",sub)
	fmt.Println("The mul of this number is this :",mul)
	fmt.Println("The div of this number is this :",div)
}

func sayHello () string{
	return "Hello, World!"
}


func add(n1, n2 int) int {
	return n1 + n2
}

func table (age , hight int , name string) string{
	return fmt.Sprintf("The age of %s is %d and the hight is %d",name,age,hight)
}

func calculater (n1,n2 int )(int,int,int,int){
	sum := n1 + n2
	sub := n1 - n2
	mul := n1 * n2
	div := n1 / n2
	return sum,sub,mul,div
}