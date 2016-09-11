package main

import (
  "os"
  "fmt"
  "strconv"
)

func fib ( n int64 ) ( ret int64 ) {
  switch( n ) {
      case 0  : ret = 0;
      case 1  : ret = 1;
      default : ret = fib( n - 1 ) + fib( n - 2 );
  }

  return;
}

func main() {
  if( len( os.Args ) < 2 ) {
    fmt.Println( "Please input an integer representing the nth fibonacci number" );
  } else {
    n, err := strconv.ParseInt( os.Args[ 1 ], 10, 64 );

    if( err != nil ) {
      fmt.Println( "Invalid input -", os.Args[ 1 ], "is not an integer" );
    } else {
      fmt.Println( fib( n ) );
    }
  }
}
