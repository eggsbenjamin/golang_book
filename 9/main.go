package main

import (
  "fmt"
  "math"
)

type Circle struct {
  radius        float64
}

func ( c *Circle ) area() float64 {
  return math.Pi * ( c.radius * c.radius );
}

func ( c *Circle ) circumference() float64 {
  return 2 * math.Pi * c.radius;
}

func ( c *Circle ) diameter() float64 {
  return 2 * c.radius;
}

func ( c *Circle ) setRadius( radius float64 ) {
  c.radius = radius;
}

func main() {
  c := Circle { 10 };
  c.setRadius( 15 );

  fmt.Println( "area :", c.area() );
  fmt.Println( "diameter :", c.diameter() );
  fmt.Println( "circumference :", c.circumference() );
}
