package main

import (
  "fmt"
  "time"
)


func sender( c chan<- string, msg string, delay int64 ) {
  for {
    c <- msg;
    time.Sleep( time.Second * time.Duration( delay ) );
  }
}

func main() {
  c1 := make( chan string );
  c2 := make( chan string );

  go sender( c1, "channel 1", 2 );
  go sender( c2, "channel 2", 5 );

  for {
    select {
      case msg1 := <- c1 :
        fmt.Println( msg1 );
      case msg2 := <- c2 :
        fmt.Println( msg2 );
    }
  }

  var input string;
  fmt.Scanln( &input );
}
