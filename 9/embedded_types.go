package main

import "fmt"

type Person struct {
  name string
}

func ( p *Person ) Talk( speech string ) {
  fmt.Println( speech );
}

type Android struct {
  Person
  serialNumber string
}

func main() {
  a := Android { serialNumber : "190586" };

  a.Person.Talk( "Hello Go!");
  a.Talk( "Hello Go!" );
}
