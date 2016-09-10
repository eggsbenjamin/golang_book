package main

import "fmt"

func main() {
  grades := [ 5 ] float32 { 
    61, 35, 94, 79, 88,
  };

  var total float32 = 0;

  for _, value := range grades {
    total += value;
  }
  
  fmt.Println( total / float32( len( grades ) ) );
}
