package main

func main() {

	namechan := make(chan string)
	agechan := make(chan int)

	go func() {
		namechan <- "Debjyoti"
	}()

	go func() {
		agechan <- 21
	}()
	for i := 0; i < 2; i++ {
		select {
		case name := <-namechan:
			println("Name is :", name)
		case age := <-agechan:
			println("Age is :", age)
		}
	}

}
