# Go Variables

<b> A typical program uses various values that may change during its execution. For Example, a program that performs some operations on the values entered by the user. The values entered by one user may differ from those entered by another user. Hence this makes it necessary to use variables as another user may not use the same values. When a user enters a new value that will be used in the process of operation, can store temporarily in the Random Access Memory of the computer and these values in this part of memory vary throughout the execution and hence another term for this came which is known as Variables. </b>

<I>Rules for Naming Variables: </I>

- Variable names must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.

```
shivang, Shivang, _shivang23  // valid variable
123Shivang, 23shivang      // invalid variable
```

- A variable name should not start with a digit.

```
234shivang  // illegal variable 
```

- The name of the variable is case sensitive.

```
shivang and Shivang are two different variables
```

- Keywords is not allowed to use as a variable name.

- There is no limit on the length of the name of the variable, but it is advisable to use an optimum length of 4 – 15 letters only.

### Declaring a Variable

<b>1. Using var keyword:</b> In Go language, variables are created using var keyword of a particular type, connected with name and provide its initial value.

Syntax:

```
var variable_name type = expression
```

Important Points:

- In the above syntax, either type or = expression can be omitted, but not both.

- If the = expression is omitted, then the variable value is determined by its type’s default value. The default value is usually 0.

Example:

```
// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

// Variable declared and
// initialized without the
// explicit type
var myvariable1 = 20
var myvariable2 = "GeeksforGeeks"
var myvariable3 = 34.80

// Display the value and the
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
								myvariable1)
									
fmt.Printf("The type of myvariable1 is : %T\n",
								myvariable1)
	
fmt.Printf("The value of myvariable2 is : %s\n",
									myvariable2)
										
fmt.Printf("The type of myvariable2 is : %T\n",
								myvariable2)
	
fmt.Printf("The value of myvariable3 is : %f\n",
									myvariable3)
										
fmt.Printf("The type of myvariable3 is : %T\n",
								myvariable3)
	
}
```

Output:

```
The value of myvariable1 is : 20
The type of myvariable1 is : int
The value of myvariable2 is : GeeksforGeeks
The type of myvariable2 is : string
The value of myvariable3 is : 34.800000
The type of myvariable3 is : float64
```

- If the expression is removed, then the variable holds zero-value for the type like zero for the number, false for Booleans, “” for strings, and nil for interface and reference type. So, there is no such concept of an uninitialized variable in Go language.

Example:

```
// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

	// Variable declared and
	// initialized without expression
	var myvariable1 int
	var myvariable2 string
	var myvariable3 float64

	// Display the zero-value of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
									myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
									myvariable2)

	fmt.Printf("The value of myvariable3 is : %f",
									myvariable3)
}
```

Output:

```
The value of myvariable1 is : 0
The value of myvariable2 is : 
The value of myvariable3 is : 0.000000
```

- If you use type, then you are allowed to declare multiple variables of the same type in the single declaration.

Example:

```
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

	// Multiple variables of the same type
	// are declared and initialized
	// in the single line
	var myvariable1, myvariable2, myvariable3 int = 2, 454, 67

// Display the values of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
									myvariable1)

fmt.Printf("The value of myvariable2 is : %d\n",
									myvariable2)

fmt.Printf("The value of myvariable3 is : %d",
									myvariable3)
}
```

Output:

```
The value of myvariable1 is : 2
The value of myvariable2 is : 454
The value of myvariable3 is : 67
```

- If you remove type, then you are allowed to declare multiple variables of a different type in the single declaration. The type of variables is determined by the initialized values.

Example:

```
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Multiple variables of different types
// are declared and initialized in the single line
var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

// Display the value and
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
									myvariable1)

fmt.Printf("The type of myvariable1 is : %T\n",
								myvariable1)

fmt.Printf("\nThe value of myvariable2 is : %s\n",
									myvariable2)

fmt.Printf("The type of myvariable2 is : %T\n",
								myvariable2)

fmt.Printf("\nThe value of myvariable3 is : %f\n",
									myvariable3)

fmt.Printf("The type of myvariable3 is : %T\n",
								myvariable3)
}
```

Output:

```
The value of myvariable1 is : 2
The type of myvariable1 is : int

The value of myvariable2 is : GFG
The type of myvariable2 is : string

The value of myvariable3 is : 67.560000
The type of myvariable3 is : float64
```

<b>2. Using short variable declaration:</b> The local variables which are declared and initialize in the functions are declared by using short variable declaration.

Syntax:

```
variable_name:= expression
```

Important Points:

- In the above expression, the type of the variable is determined by the type of the expression.

Example:

```
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
myvar1 := 39
myvar2 := "GeeksforGeeks"
myvar3 := 34.67

// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}
```

Output:

```
The value of myvar1 is : 39
The type of myvar1 is : int

The value of myvar2 is : GeeksforGeeks
The type of myvar2 is : string

The value of myvar3 is : 34.670000
The type of myvar3 is : float64
```

- Most of the local variables are declared and initialized by using short variable declarations due to their brevity and flexibility.

- The var declaration of variables are used for those local variables which need an explicit type that differs from the initializer expression, or for those variables whose values are assigned later and the initialized value is unimportant.

- Using short variable declaration you are allowed to declare multiple variables in the single declaration.

Example:

```
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Multiple variables of same types
// are declared and initialized in
// the single line
myvar1, myvar2, myvar3 := 800, 34, 56

// Display the value and
// type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}
```

Output:

```
The value of myvar1 is : 800
The type of myvar1 is : int

The value of myvar2 is : 34
The type of myvar2 is : int

The value of myvar3 is : 56
The type of myvar3 is : int
```

-Using short variable declaration you are allowed to declare multiple variables of different types in the single declaration. The type of these variables are determined by the expression.

Example:

```
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Using short variable declaration
// Multiple variables of different types
// are declared and initialized in the single line
myvar1, myvar2, myvar3 := 800, "Geeks", 47.56

// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)

fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)

fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)

}
```

Output:

```
The value of myvar1 is : 800
The type of myvar1 is : int

The value of myvar2 is : Geeks
The type of myvar2 is : string

The value of myvar3 is : 47.560000
The type of myvar3 is : float64
```
