#  Arrays in Go

Arrays in Golang or Go programming language is much similar to other programming languages. In the program, sometimes we need to store a collection of data of the same type, like a list of student marks. Such type of collection is stored in a program using an Array. An array is a fixed-length sequence that is used to store homogeneous elements in the memory. Due to their fixed length array are not much popular like Slice in Go language. In an array, you are allowed to store zero or more than zero elements in it. The elements of the array are indexed by using the [] index operator with their zero-based position, which means the index of the first element is array[0] and the index of the last element is array[len(array)-1]. 

<p align = "center">
<img src = "https://user-images.githubusercontent.com/101946115/211153393-d3fa4d71-c431-4d21-89fa-0fe39cec9f87.png" />
  </p>
  
## Creating and accessing an Array

In Go language, arrays are created in two different ways:

Using var keyword: In Go language, an array is created using the var keyword of a particular type with name, size, and elements. 

Syntax:

```
Var array_name[length]Type
```

Important Points:

In Go language, arrays are mutable, so that you can use array[index] syntax to the left-hand side of the assignment to set the elements of the array at the given index.

```
Var array_name[index] = element
```

- You can access the elements of the array by using the index value or by using for loop.
- In Go language, the array type is one-dimensional.
- The length of the array is fixed and unchangeable.
- You are allowed to store duplicate elements in an array.


