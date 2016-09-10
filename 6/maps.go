package main

import "fmt"

func main() {
  hashTable := make( map [ string ] int );

  hashTable[ "key1" ] = 1;
  hashTable[ "key2" ] = 2;
  
  delete( hashTable, "key2" );

  fmt.Println( hashTable[ "key1" ] );

  fmt.Println( hashTable[ "key2" ] );

  //  elements...

  elements := map [ string ] map [ string ] string {
    "H": map [ string ] string {
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map [ string ] string {
      "name":"Helium",
      "state":"gas",
    },
    "Li": map [ string ] string {
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map [ string ] string {
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map [ string ] string {
      "name":"Boron",
      "state":"solid",
    },
    "C":  map [ string ] string {
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map [ string ] string {
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map [ string ] string {
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map [ string ] string {
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map [ string ] string {
      "name":"Neon",
      "state":"gas",
    },
  }

  if elem, ok := elements[ "H" ]; ok {
    fmt.Println( elem, ok );
  }
}
