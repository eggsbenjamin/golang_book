#Functions

A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. 
- Golang functions can have more than one return value

Functions (also known as procedures or subroutines) are often represented as a black box
- i.e. you only care about the input and output

##Anatomy Of A Function

Functions are...
  - declared using the `func` keyword
  - followed by the name of the function
  - followed by the parameters contained in parenthesis declared ( name type, name, type ...etc ) 
  - followed by the return type(s)
  - followed by the function body contained in curly braces

e.g.

```
func myFunction ( myParam1 string, myParam2 int ) string {
  //  do stuff...
}
```

Collectively the parameters and the return type are called the function's signature

##Return Type(s)

Return types can be declared... 
  - as a single unamed type e.g. `func myFunc ( param string ) string` || `func myFunc ( param string ) ( string )`
  - as multiple unamed types e.g. `func myFunc ( param string ) ( string, int )`
  - or, as above, but using named types e.g. `func myFunc ( param string ) ( namedType string )` || `func myFunc ( param string ) ( namedType1 string, namedType2 int )`
    - named return types must by declared within parenthesis

You cannot declare a mixture of named and unamed return types. They must all be named or unamed.

The most common pattern for multiple return types is for the final return type to be a boolean indicating the success of the operation.

e.g.

```
func shortenBenjamin ( longName string ) ( string, bool ) {
  if( longName != "Benjamin" ) {
    return "", false;
  } else {
    return "Ben", true;
  }
}   
``` 

##Named Return Types

When named return type(s) are declared as a part of a function's signature...
  - they can't be re-declared within the function body
  - the return statement can be declared without without a value resulting in the variables corresponding to the named return types being returned in the order that the were declared.

##Variadic Functions

The last parameter of a function can take a special form e.g. `func myFunc ( param1 string, param2 ...int )`

When the last parameter of a function is declared using this syntax zero or more of the type can be passed as arguments to the function 

e.g

```
func myFunc ( args... int ) {
  for _, value := range args {
    fmt.Println( value );
  } 
}

myFunc( 1, 2, 3, 4, 5 ); 

/*  
  outputs :

  1
  2
  3
  4
  5
*/
```

##Closure

So far each function that has been described has been a declared using a function expression statement. 

You can also declare functions using a function assignment expression.

e.g.

```
funcVar := func( param int ) int {
  return param + 5;
}

funcVar( 5 );

// returns 10
```

You can create functions within functions by declaring the inner function(s) using the function assignment syntax.

Functions declared within another function have access to their parent function's scope ( but not vice-versa ).

A closure is a function that access members outside of it's scope.

e.g.

```
func parentFunc ( parentParam int ) int {
  childFunc := func( childParam int ) int {
    return parentParam + childParam;
  }

  return childFunc( 5 );
}

parentFunc( 10 );

//  returns 15
```

##Recursion

Functions are able to call themselves

e.g.

```
func myFunc( param uint ) string {
  if( param == 0 ) {
    return "zero";
  }

  fmt.Println( param );

  param--;

  return myFunc( param );
}

myFunc( 5 );

/*  
  outputs :

  5
  4
  3
  2
  1
  zero
*/
```

##Defer, Panic and Recover

###Defer

A defer statement schedules a function to be called after the function in which it is declared completes

Deferred functions are are called even if a runtime panic occurs making them ideal for freeing resources etc...

e.g.

```
func myFunc() {
  fmt.Println( "myFunc called!" );
}

func anotherFunc() {
  defer myFunc();
  
  panic(`
    I really hope that 'myFunc' was called...
  `);
}

/*
  outputs :

  I really hope that 'myFunc' was called...
*/
```

###Panic & Recover

A panic indicates a runtime error and can...
  - happen unexpectedly
  - be called explicitly e.g. `panic( "Oh No!" );` ( see above )

the recover function handles a runtime panic by...
  - stopping the runtime panic
  - returning the value passed to panic

A panic will stop execution of a function immediately so to 'catch' a panic you must pair it with the the defer function

```
func myFunc() {
  defer func() {
    str := recover();
    fmt.Println( "Recovered :", str );
  }();

  panic( "Something's Gone Horribly Wrong!!!" );
}

/*
  outputs :

  Recovered : Something's Gone Horribly Wrong!!!  
*/
```

#Problems

##sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?

- `func sum ( numbers [] float64 ) float64 { } ` 
- the signature of a function is made up of it's paramter(s) and return type(s) 

##Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

- see half.go

##Write a function with one variadic parameter that finds the greatest number in a list of numbers.

- see greatest.go

##Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

- see makeOddGenerator.go

##The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

- see fibonacci.go

##What are defer, panic and recover? How do you recover from a run-time panic?

- panic
  - a function that explicitly throw a run-time panic with the provided argument as a message
- recover 
  - a function that recovers from a run-time panic and returns the panic message
  - must be used in conjuntion with defer i.e. recover must be called within a deferred function to 'catch' the panic
- defer 
  - a statement that must be declared within a function
  - allows you to specify a function to be called after the containing function has completed  
  - the specified function is called regardless of the outcome of the containing function i.e. success/panic
