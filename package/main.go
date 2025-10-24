package main

import (
	"fmt"

	"github.com/Debjyoti2004/ToDo/auth"
)

func main() {
	fmt.Println(auth.SignIn("user", "pass"))
	fmt.Println(auth.SignUp("user", "pass", "debjyotishit27@gmail.com"))
}

// Commands for Go module management and execution:
//---------------------------------------------------------
// go mod init github.com/Debjyoti2004/ToDo // For initializing module
// go mod add github.com/Debjyoti2004/ToDo/auth // For adding local package
// go mod tidy // For cleaning up unused dependencies
// go run package/main.go // For running the main file
// go build package/main.go // For building the main file executable
// go clean -cache // For cleaning the module cache
//---------------------------------------------------------
