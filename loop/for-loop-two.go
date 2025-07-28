package main
import "fmt"

func main(){
	animals := make(map[string]string)
	animals["Dog"] = "Woof"
	animals["Cat"] = "Meow"
	animals["Bird"] = "Tweet"

	for key, value := range animals{
		fmt.Println("The value of animal is:",value,"and the key is:",key)
	}
}