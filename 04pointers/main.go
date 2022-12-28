package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)
}
