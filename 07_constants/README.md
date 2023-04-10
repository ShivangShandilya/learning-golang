# Constants

As the name CONSTANTS suggests, it means fixed. In programming languages also it is same i.e., once the value of constant is defined, it cannot be modified further. There can be any basic data type of constants like an integer constant, a floating constant, a character constant, or a string literal. 

### Untyped and Typed Numeric Constants: 

Typed constants work like immutable variables can inter-operate only with the same type and untyped constants work like literals can inter-operate with similar types. Constants can be declared with or without a type in Go. Following is the example which show typed and untyped numeric constants that are both named and unnamed. 

```
const untypedInteger          = 123
const untypedFloating          = 123.12

const typedInteger  int             = 123
const typedFloatingPoint   float64  = 123.12
```

Following is a list of constants in Go Language: 

- Numeric Constant (Integer constant, Floating constant, Complex constant)
- String literals
- Boolean Constant

Numeric Constant: Numeric constants are high-precision values. As Go is a statically typed language that does not allow operations that mix numeric types. You can’t add a float64 to an int, or even an int32 to an int. Although, it is legal to write 1e6*time.Second or math.Exp(1) or even 1<<(‘\t’+2.0). In Go, constants, unlike variables, behave like regular numbers. 

Numeric constant can be of 3 kinds:

1. integer
2. floating-point
3. complex 

### Integer Constant: 

- A prefix specifies the base or radix: 0x or 0X for hexadecimal, 0 for octal, and nothing for decimal.
- An integer literal can also have a suffix that is a combination of U(upper case) and L(lower case), for unsigned and long, respectively
- It can be a decimal, octal, or hexadecimal constant.
- An int can store at maximum a 64-bit integer, and sometimes less.

### Complex constant: 

Complex constants behave a lot like floating-point constants. It is an ordered pair or real pair of integer constant( or parameter). And the constant are separated by a comma, and the pair is enclosed in between parentheses. The first constant is the real part, and the second is the imaginary part. A complex constant, COMPLEX*8, uses 8 bytes of storage. 

### Floating Type Constant: 

- A floating type constant has an integer part, a decimal point, a fractional part, and an exponent part.
- Can be represented as a floating constant either in decimal form or exponential form.
- While represented using the decimal form, it must include the decimal point, the exponent, or both.
- And while represented using the exponential form, it must include the integer part, the fractional part, or both.

Following are examples of Floating type constants: 

```
3.14159       /* Legal */
314159E-5L    /* Legal */
510E          /* Illegal: incomplete exponent */
210f          /* Illegal: no decimal or exponent */
.e55          /* Illegal: missing integer or fraction */
```
