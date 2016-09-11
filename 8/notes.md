#Pointers

Pointers contain a reference to the location in memory where a value is stored rather than the value itself.

e.g.

```
func myFunc( ptr *int ) {
  fmt.Println( "memory address :", &ptr );
  fmt.Println( "value :", *ptr );
}

x := 5;

myFunc( &5 );

/*
  output :

  memory address : 0x4201a0188
  value : 5 
*/
```

##The & and * Operators

In Golang a pointer parameter is declared by using the `*` operator before the type

e.g.

```
func myFunc ( pointer *int ) {...
```

The `*` operator is also used to 'dereference' pointer variables.

'Dereferencing' a pointer variable is to...
  - access the value located at the memory address
  - or modify the value located at the memory address

The `&` operator is used to find the memory address of a variable.
 
e.g.

```
func myFunc ( pointer *int ) {
  fmt.Println( "memory address :", &pointer );
  fmt.Println( "value :", *pointer );

  *pointer = 20;

  fmt.Println( "memory address :", &pointer );
  fmt.Println( "value :", *pointer );
}

x := 10;

myFunc( &x ); //  pass the memory address of the variable x to myFunc...

/*
  outputs :
  
  memory address : 0x4201a0188
  value : 5 
  memory address : 0x4201a0188
  value : 20 
*/
```

##new

`new` is a built-in function that...
  - takes a type as an argument
  - allocates a space in memory to hold the type e.g. 64 bits for an int64
  - returns a pointer to the memory address

e.g

```
x := new( int );

fmt.Println( x );

/*
  outputs :

  0x4201a4000
*/
```

#Problems

##How do you get the memory address of a variable?

- by prefixing a variable name with the `&` operator 

##How do you assign a value to a pointer?

- by prefixing the pointer variable name with the `*` operator and using it in an assignment statement

##How do you create a new pointer?

- by passing the type, that you want to create a pointer to, to the `new` function 

##What is the value of x after running this program:

```
func square(x *float64) {
  *x = *x * *x
}

func main() {
  x := 1.5
  square(&x)
}
```

- x == 0.75 

##Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

- see swap.go
