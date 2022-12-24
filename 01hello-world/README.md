# Hello World in Golang

Hello, World! is the first basic program in any programming language. Let’s write the first program in the Go Language using the following steps:

- First of all open Go compiler. In Go language, the program is saved with .go extension and it is a UTF-8 text file.
- Now, first add the package main in your program:

```
package main
```

Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code. In Go, there are two types of program available one is executable and another one is the library.

- After adding main package import “fmt” package in your program:

```
import(
"fmt"
)
```

- Now write the code in the main function to print hello world in Go language:

```
func main() {
    fmt.Println("!... Hello World ...!")
}
```

In the above lines of code we have:

- func: It is used to create a function in Go language.

- main: It is the main function in Go language, which doesn’t contain the parameter, doesn’t return anything, and call when you execute your program.

- Println(): This method is present in fmt package and it is used to display “!… Hello World …!” string. Here, Println means Print line.
