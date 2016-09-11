package main

import "fmt"

func printAddressAndValue( ptr *int ) {
  fmt.Println( "memory address :", &ptr );
  fmt.Println( "value :", *ptr );

  *ptr = 20;

  fmt.Println( "memory address :", &ptr );
  fmt.Println( "value :", *ptr );
}

func main() {
  x := new( int );

  printAddressAndValue( x );
}
