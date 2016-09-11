package main

import "fmt"

func parentFunc ( parentParam int ) int {
  childFunc := func( childParam int ) int {
    return parentParam + childParam;
  }

  return childFunc( 5 );
}

func main() {
  fmt.Println( parentFunc( 10 ) );
}
