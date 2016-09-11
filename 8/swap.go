package main

import "fmt"

func swap( x *int, y *int ) {
  z := *x; //  hold the value of x...

  *x = *y;
  *y = z;
}

func main() {
  x := 1;
  y := 2;

  fmt.Println( "x =", x );
  fmt.Println( "y =", y );

  swap( &x, &y );

  fmt.Println( "x =", x );
  fmt.Println( "y =", y );
}
