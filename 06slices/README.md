# Slices in Golang

In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure. Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice. It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array. Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. It is allowed to store duplicate elements in the slice. The first index position in a slice is always 0 and the last one will be (length of slice – 1).

## Declaration of Slice

A slice is declared just like an array, but it doesn’t contain the size of the slice. So it can grow or shrink according to the requirement. 

Syntax:  

```
[]T
or 
[]T{}
or 
[]T{value1, value2, value3, ...value n}
```

Here, T is the type of the elements. For example: 

```
var my_slice[]int 
```

## Components of Slice

A slice contains three components:

- Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. Here, it is not necessary that the pointed element is the first element of the array.
- Length: The length is the total number of elements present in the array.
- Capacity: The capacity represents the maximum size upto which it can expand.
