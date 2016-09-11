package main

import "fmt"

func half( x int ) ( y int, even bool ) {
  if ( x % 2 == 0 ) {
    y = x / 2;
    even = true;
  } else {
    y = 0;
    even = false;
  }

  return;
}

func main() {
  values := [] int { 1, 2, 3, 4, 5, 6, 7, 8, 9, };

  for _, value := range values {
    result, even := half( value );
    fmt.Println( value, ":", "Even =", even, ",Half =", result );
  }
}
