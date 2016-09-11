package main

import "fmt"

func average( values [] float64 ) ( avg float64 ) {
  for _, value := range values {
    avg += value;
  }

  avg /= float64( len( values ) );

  return;
}

func main() {
  fmt.Println( average( [] float64 { 1, 2, 3, } ) );
}
