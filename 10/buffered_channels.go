package main

import (
  "fmt"
  "time"
)

func sender( c chan<- string, msg string, delay int64 ) {
  for {
    time.Sleep( time.Second * time.Duration( delay ) );
    c <- msg;
  }
}

func main() {
  c := make( chan string, 1 );

  go sender( c, "1", 3 );
  go sender( c, "2", 6 );
  go sender( c, "3", 9 );
  go sender( c, "4", 12 );
  go sender( c, "5", 15 );

  for i := 0; i < 5; i++ {
    var input string;

    fmt.Scanln( &input );
    fmt.Println( <- c );
  }
}
