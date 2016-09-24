package main

import (
  "fmt"
  "time"
)

//  block execution with channel

func Sleep( delay float64 ) {
  result := <-time.After( time.Duration( delay ) * time.Second );
  _ = result;
}

func main() {
  fmt.Println( "I'm tired..." );

  Sleep( 10 );

  fmt.Println( "I feel energised!!!" );
}
