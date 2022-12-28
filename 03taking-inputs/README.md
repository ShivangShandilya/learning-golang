# How to Take Input from the User in Golang?

Scanln function can be used to take the input from the user in the Golang. Below is the example of taking input from the user:

```
// Golang program to show how
// to take input from the user
package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Name: ")

	// var then variable name then variable type
	var first string

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Last Name: ")
	var second string
	fmt.Scanln(&second)

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")

	// Addition of two string
	fmt.Print(first + " " + second)
}

```
Now save this file and execute as shown in the below screenshot:

![image](https://user-images.githubusercontent.com/101946115/209806991-da5bbe75-d0f7-4e1e-95d1-5f4d66c3edeb.png)


## Description About the Program:

- <b>main Package:</b> When we build reusable pieces of code, we will develop a package as a shared library. But when we develop an executable program, we will use the package “main” for making the package as an executable program. The package “main” tells the Golang compiler that the package should compile as an executable program instead of a shared library. The main function in the package “main” will be the entry point of our executable program. Remember when we build shared libraries, we will not have any main package and main function in the package.

- <b>fmt.Println</b> is the print function which is used to print the output in the next line. While fmt.Print is used to display output in the same line. Whatever needs to be printed has to be written in inverted commas ” “.

- <b>var first string</b> is the declaration of the variable first which is of string type. To declare variables following syntax needs to be followed:

```
 var var_name data_type 
```

- <b>fmt.Scanln(&first)</b> is used to take input from user in the next line. While fmt.Scan is used to take input from user in the same line. Ampersand is necessary to give the reference as to in which variable we have to store the variable.

- The Last line will simply add the two string of First Name and Last Name and will output the Full Name.
