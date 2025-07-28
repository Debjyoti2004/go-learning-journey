package main

import "fmt"

func printSlice[T any](item []T) {

	for _, item := range item {
		fmt.Println(item)
	}

}

func main() {
	printSlice([]string{"a", "b", "c"})
	printSlice([]int{1, 2, 3})
}
