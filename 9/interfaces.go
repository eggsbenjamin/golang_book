package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area() float64
}

type Circle struct {
  Shape
  radius float64
}

func ( c *Circle ) area() float64 {
  return math.Pi * ( c.radius * c.radius );
}

func main() {
  c := Circle { radius : 10 };

  fmt.Println( c.area() );
}
