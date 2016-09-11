package main

import (
  "fmt"
  "math"
)

type Shape interface {
  area()      float64
  perimeter() float64
}

type Rectangle struct {
  Shape
  shortSideLength float64
  longSideLength  float64
}

type Circle struct {
  Shape
  radius float64
}

func ( r *Rectangle ) area() float64 {
  return ( r.shortSideLength * 2 ) + ( r.longSideLength * 2 );
}

func ( c *Circle ) area() float64 {
  return math.Pi * ( c.radius * c.radius );
}

func ( r *Rectangle ) perimeter() float64 {
  return r.shortSideLength * r.longSideLength;
}

func ( c *Circle ) perimeter() float64 {
  return 2 * math.Pi * c.radius;
}

func main() {
  c := Circle    { radius : 10 };
  r := Rectangle { shortSideLength : 10, longSideLength : 15 };

  fmt.Println( "Circle - Area :", c.area(), "Circumference :", c.perimeter() );
  fmt.Println( "Rectangle - Area :", r.area(), "Perimeter :", r.perimeter() );
}
