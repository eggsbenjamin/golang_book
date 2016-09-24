package main

import "fmt"

func main() {
  var input string;

  for i := 0; i < 100; i++ {
    var r := go fmt.Println( i );
    fmt.Println( r );
  }

  fmt.Scanln( &input );
}
