package main

import "fmt"

func first() {
  fmt.Println( "first" );
}

func second() {
  fmt.Println( "second" );
}

func main() {
  //  call this after the main function has completed regardless 
  //  of the outcome...
  defer second(); 
  first();

  panic( `
    I really hope that the outputs from both functions 
    are printed above?!!
  `);
}
