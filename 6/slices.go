package main

import "fmt"

func main() {
  underlyingArray := [ 10 ] float64 {
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
  };

  slice1 := underlyingArray[ 3: ];

  slice2 := append( slice1, 11, 12, 13 );

  slice3 := make( [] float64, 2 );

  copy( slice3, slice2 ); //  copy slice2 to slice3...

  fmt.Println( slice1 );

  fmt.Println( slice2 );

  fmt.Println( slice3 );
}
