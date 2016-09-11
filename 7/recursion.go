package main

import "fmt"

func recursiveCountdown( x uint ) string {
  if( x == 0 ) {
    return "zero";
  }

  fmt.Println( x );

  x--;

  return recursiveCountdown ( x );
}

func main() {
  fmt.Println( recursiveCountdown( 10 ) );
}
