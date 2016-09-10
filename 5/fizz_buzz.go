package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    var divisibleByThree bool = ( i % 3 == 0 );
    var divisibleByFive  bool = ( i % 5 == 0 );

    if( divisibleByThree && divisibleByFive ) {
      fmt.Println( "FizzBuzz" );
    } else if ( divisibleByThree ) {
      fmt.Println( "Fizz" );
    } else if ( divisibleByFive ) {
      fmt.Println( "Buzz" );
    } else {
      fmt.Println( i );
    }
  }
}
