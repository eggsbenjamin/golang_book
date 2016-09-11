package main

import "fmt"

func printAllInts( args ...int ) {
  for _, value := range args {
    fmt.Println( value );
  } 
}

func main() {
  //  you can pass as many ints as you like to printAllInts!!!

  printAllInts( 1, 2, 3, 4, 5, 6, 7, 8, 9 , 10000 );

  //  here we pass an array of ints to the function but use the '...' suffix
  //  to pass the values as individual integers per the function's signature

  printAllInts( [] int { 1, 7, 8, 9, 10, 12, 13, 14, }... ); 
}
