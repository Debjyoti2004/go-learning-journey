package main
import "fmt"

type Animal interface{
	Speak() string
	NumberOfLegs() int
}

type Dog struct{
	Name string
	Breed string
}

type Cat struct{
	Name string
	Breed string
}

func (d *Dog) Speak() string{
	return "Woof"
}

func (d *Dog) NumberOfLegs() int{
	return 4
}

func (c *Cat) Speak() string{
	return "Meow"
}

func (c *Cat) NumberOfLegs() int{
	return 4
}

func main(){
	dog := Dog{Name: "Rex", Breed: "German Shepherd"}
	cat := Cat{Name: "Whiskers", Breed: "Siamese"}

	animals := []Animal{&dog, &cat}

	for _, animal := range animals{
		fmt.Println(animal.Speak()," ",animal.NumberOfLegs())
	}
}
