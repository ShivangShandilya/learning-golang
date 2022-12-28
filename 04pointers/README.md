# Pointers in Golang

Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable. Pointers in Golang is also termed as the special variables. The variables are used to store some data at a particular memory address in the system. The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

### What is the need for the pointers? 

To understand this need, first, we have to understand the concept of variables. Variables are the names given to a memory location where the actual data is stored. To access the stored data we need the address of that particular memory location. To remember all the memory addresses(Hexadecimal Format) manually is an overhead that’s why we use variables to store data and variables can be accessed just by using their name. 
Golang also allows saving a hexadecimal number into a variable using the literal expression i.e. number starting from 0x is a hexadecimal number.


A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points where the memory is located and provides ways to find out the value stored at that memory location. It is generally termed as a Special kind of Variable because it is almost declared as a variable but with *(dereferencing operator).

<p align = "center">
<img src ="https://user-images.githubusercontent.com/101946115/209807292-10bb4532-66b6-464e-baf3-46d91ac7c020.png" height = 400 width = 700 />
  </p>
  
  
## Declaration and Initialization of Pointers

Before we start there are two important operators which we will use in pointers i.e. 

- <b>* Operator</b> also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.

- <b>& operator</b> termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

### Declaring a pointer:

```
var pointer_name *Data_Type
```

Example: Below is a pointer of type string which can store only the memory addresses of string variables.

```
var s *string
```

### Initialization of Pointer: 

To do this you need to initialize a pointer with the memory address of another variable using the address operator as shown in the below example:

```
// normal variable declaration
var a = 45

// Initialization of pointer s with 
// memory address of variable a
var s *int = &a
```

Example:

```
// Golang program to demonstrate the declaration
// and initialization of pointers
package main

import "fmt"

func main() {

	// taking a normal variable
	var x int = 5748
	
	// declaration of pointer
	var p *int
	
	// initialization of pointer
	p = &x

		// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
}
```

Output: 

```
Value stored in x =  5748
Address of x =  0x414020
Value stored in variable p =  0x414020
```


