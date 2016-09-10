package main

import "fmt"

func main() {
  var fah float64;

  fmt.Print( "Enter a Fahrenheit value to be converted to Celcius\n" );
  fmt.Scanf( "%f", &fah );

  var cel float64 = ( fah - 32 ) * 5/9;

  fmt.Println( fah, "F ==", cel, "C" );
}
