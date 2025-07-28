package main
import "fmt"

func main(){
	number := []int{1,2,3,4,5,6,7,8,9,10}

	number = append(number, 11,12,13,14,15)
	fmt.Println(number)

	fmt.Println(number[0:3])
	fmt.Println(number[3:])
	fmt.Println(number[:3])
	fmt.Println(number[:])
	fmt.Println(number[3:6])
	fmt.Println(number[6:9])
	fmt.Println(number[9:12])
	fmt.Println(number[12:15])
}