#Structs and Interfaces

##Structs

A struct is a type that contains named fields

e.g.

```
type myStruct struct {
  field1 string
  field2 int64
  field3 float64
}
```

In the above example we can see...
  - the `type` keyword introducing a new type
  - the name, `myStruct`, of the struct being defined
  - the `struct` keyword indicating that we are defining a struct
  - a list of the fields to be contained within the type
    - only delimeted by commas if they're declared on a single line

###Initialization

A struct can be initialized in multiple ways

e.g.

```
//  create a new instance of myStruct and assign it to the variable 'emptyStruct'
//  all fields of emptyStruct will be set to their default values e.g. 0 for float64

var emptyStruct myStruct; 


//  the same as above but with a subtle difference
//  the new function returns a pointer to rather than the value

emptyStruct := new( myStruct );

//  create a new instance of myStruct and assign it to the variable 'populatedStruct'
//  each field within 'populatedStruct' is assigned a value using this syntax

populatedStruct := myStruct { field1 : "hello", field2 : 4, field3 : 53.8 };

//  the same as above but when you don't supply the names of the fields
//  they're populated in the order in which they're defined

populatedStruct := myStruct { "hello", 4, 53.8 };
```

###Fields

You can access the fields within a struct usung the `.` operator.

e.g.

```
myStruct.field1 := "World";

fmt.Println( myStruct.field1 );

/*  
  outputs :

  World
*/
```

When declaring a function that takes struct(s) as arguments...
  - it's best to declare the parameter as a pointer to a struct
  - if a pointer is passed you can modify the fields
  - you can still access fields using the `.` operator when using a pointer  

e.g.

```
func myFunc ( m *myStruct ) int {
  //  access/modify fields even though it's a pointer 
  m.field1 = "Hello World";
  
  return 1; 
}
```

###Methods

In Golang a method is a function that belongs to a struct.

Methods are defined much like functions but with a subtle difference.

e.g.

```
func ( m *myFunc ) concat() string {
  return fmt.Println( m.field1, m.field2, m.field3, m.field4 );
}
```

The declaration above shows...
  - the first set of parenthesis defines the type the method belongs to and the name of the internal pointer variable
  - the name of the method 'concat' in this instance ( you can declare parameters here also )
  - the return type(s)

###Embedded Types

You can declare a struct as a field within another struct 

e.g.

```
type Person struct {
  name string
}

func ( p *Person ) Talk( speech string ) {
  fmt.Println( speech );
}

type Android struct {
  Person
  serialNumber string
}

func main() {
  a := Android { serialNumber : "190586" };

  a.Talk( "Hello Go!" );
  a.Person.Talk( "Hello Go!");
}
```

You declare an embedded type by supplying it as an un-named field when declaring the struct

You can call any of the methods belonging to the embedded type with or without specifying the embedded type in the function call

Embedded types are similar to inheritance in other languages

##Interfaces

An interface defines a set of methods, a 'method set', that a struct must have to imlement the interface

Interfaces are used for defining common functionality between structs

e.g.

```
type Shape interface {
  area() float64
}

type Circle struct {
  Shape
  radius float64
}

func ( c *Circle ) area() float64 {
  return math.Pi * ( c.radius * c.radius );
}
```

In the above example we can see...
  - the definition of an interface called Shape
  - an interface's definition is similar to a struct definition 
  - interfaces can only contain methods
  - interfaces don't implement methods
  - the Circle struct implements the Shape interface

You can declare interfaces as parameters to functions/methods and pass any struct that implements that interface as an argument.

#Problems

##What's the difference between a method and a function?

- a method is a field within a struct
- the syntax for declaring a method is different from that of a method 

e.g.

```
//  function declaration

func myFunc ( x int ) int {

}

//  method declaration

func ( m *myStruct ) myFunc( x int ) int {

} 
```

##Why would you use an embedded anonymous field instead of a normal named field?

- So that you're able to call methods belonging to the embedded anonymous field without specifying a name

##Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.

- see perimeter.go
