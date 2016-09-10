package main

import "fmt"

func main() {
  var ft float64; 

  fmt.Print( "Enter a value in feet to be converted to metres\n" );
  fmt.Scanf( "%f", &ft );

  var m float64 = ft * 0.3048;

  fmt.Println( ft, "ft ==", m, "M" );
}
