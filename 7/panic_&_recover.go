package main

import "fmt"

func main() {
  defer func() {
    str := recover();
    fmt.Println( "Recovered :", str );
  }();

  panic( "Something's Gone Horribly Wrong!!!" );
}
