package main

import (
  "fmt"
  "time"
  )

func printer( c <-chan string ) {
  for {
    msg := <- c;
    fmt.Println( msg );
    time.Sleep( time.Second * 1 );
  }
}

func sendMsg( c chan<- string, msg string ) {
  for {
    c <- msg;
  }
}

func main() {
  var c chan string = make( chan string );

  go sendMsg( c, "ping" );
  go sendMsg( c, "pong" );
  go sendMsg( c, "ding" );
  go sendMsg( c, "dong" );
  go printer( c );

  var input string;
  fmt.Scanln( &input );
}
