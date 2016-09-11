package main

import "fmt"

func findGreatest( values [] float64 ) ( greatest float64 ) {
  if( len( values ) > 0 ) {
    greatest = values[ 0 ];

    for _, value := range values {
      if( value > greatest ) {
        greatest = value;
      }
    }

    return;
  } else {
    panic( "The Array must contain at least one element!" );
  }
}

func main() {
  numbers := [] float64 { 0, 77, 2, 81, 94, 43, 59 };

  fmt.Println( "Greatest :", findGreatest( numbers ) );
}
