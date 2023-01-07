package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Shareefa"
	fruitList[3] = "Mango"

	fmt.Println("The fruit list is: ", fruitList)

	var vegies = [3]string{"Spinach", "Cauliflower", "Potato"}
	fmt.Println(vegies)
}
