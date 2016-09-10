package main

import "fmt"

func main() {
  var x string = "Hello World";

  fmt.Println( x );

  var y string;
  y = "World Hello";

  fmt.Println( y );

  y = "Hello World";

  fmt.Println( y );

  y = y + " dlroW olleH";

  fmt.Println( y );

  z := "Hello Shorthand";

  fmt.Println( z );

  z += " dnahtrohS olleH";

  fmt.Println( z );
}
