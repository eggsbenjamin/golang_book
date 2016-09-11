package main

import "fmt"

func shortenBenjamin ( longName string ) ( string, bool ) {
  if( longName != "Benjamin" ) {
    return "", false;
  } else {
    return "Ben", true;
  }
}   

func main() {
  if _, ok := shortenBenjamin( "Darren" ); ok {
    fmt.Println( "Correct Name" );
  } else {
    fmt.Println( "Incorrect Name" );
  }
}
