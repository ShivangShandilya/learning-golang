# Go Programming Language

![image](https://user-images.githubusercontent.com/101946115/209450673-6392aeda-88c0-4abb-978d-559eb822f32a.png)

Golang is a procedural and statically typed programming language having the syntax similar to C programming language. Sometimes it is termed as Go Programming Language. It provides a rich standard library, garbage collection, and dynamic-typing capability. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language and mainly used in Google’s production systems. Golang is one of the most trending programming languages among developers.

##  Beginning with Go programming

There are various online IDEs such as The Go Playground, repl.it, etc. which can be used to run Go programs without installing. 

For installing Go in own PCs or Laptop we need of following two software: Text editor and Compiler 

Text Editor: Text editor gives you a platform where you write your source code. Following are the list of text editors:  

- Windows notepad
- OS Edit command
- Brief
- Epsilon
- vm or vi
- Emacs
- VS Code

Finding a Go Compiler: Go distribution comes as a binary installable for FreeBSD (release 8 and above), Linux, Mac OS X (Snow Leopard and above), and Windows operating systems with 32-bit (386) and 64-bit (amd64) x86 processor architectures.

## Why this “Go language”?

Because Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aims to be modern, with support for networked and multicore computing. 

## What excluding in Go which is present in other languages? 

- Go attempts to reduce the amount of typing in both senses of the word. Throughout its design, developers tried to reduce clutter and complexity.
- There are no forward declarations and no header files; everything is declared exactly once.
- Stuttering is reduced by simple type derivation using the := declare-and-initialize construct.
- There is no type hierarchy: types just are, they don’t have to announce their relationships.

## Hardware Limitations

We have observed that in a decade, the hardware and processing configuration is changing at a very slow rate. In 2004, P4 was having the clock speed of 3.0 GHz and now in 2018, Macbook pro has the clock speed of Approx (2.3Ghz v 2.66Ghz). To speed up, the functionality we use more processors, but using more processor the cost also increases. And due to this we use limited processors and using limited processor we have a heavy programming language whose threading takes more memory and slows down the performance of our system. Hence, to overcome such problem Golang has been designed in such a way that instead of using threading it uses Goroutine, which is similar to threading but consumes very less memory. 
Like threading consumes 1MB whereas Goroutine consumes 2KB of memory, hence at the same time, we can have millions of goroutine triggered. 
So the above-discussed point makes golang a strong language that handles concurrency like C++ and Java. 

## Advantages and Disadvantages of Go Language

Advantages:  

1. Flexible- It is concise, simple and easy to read.
2. Concurrency- It allows multiple process running simultaneously and effectively.
3. Quick Outcome- Its compilation time is very fast.
4. Library- It provides a rich standard library.
5. Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.
6. It validates for the interface and type embedding.

Disadvantages:  

1.It has no support for generics, even if there are many discussions about it.

2.The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.


