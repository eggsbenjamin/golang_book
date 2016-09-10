#Notes

##What are two ways to create a new variable?

- Using the long hand syntax e.g. [var|const] x string or...
- Using the shorthand syntax x := "Hello World"

##What is the value of x after running: x := 5; x += 1?

- x == 6

##What is scope and how do you determine the scope of a variable in Go?

- scope is the context in which variables exist/are accessible 
- You can determine the scope of a variable by figuring out which are the outer-most curly braces that contain a variable
- *NOTE - if a variable is declared in the global scope, i.e. outside of any curly braces, then it is accessible anywhere

##What is the difference between var and const?

- var creates a variable which can have different values assigned to it over the course of it's life time.
- var can also be declared without a value assigned to it
- const creates a variabled that can only have a variable assinged to it once. If it is changed an error is thrown
- const must be assigned a value at declaration

##Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

- See f2c.go

##Write another program that converts from feet into meters. (1 ft = 0.3048 m)

- See f2m.go
