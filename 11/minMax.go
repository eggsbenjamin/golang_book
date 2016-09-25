package minMax

import (
  "fmt"
  "errors"
)


//  Finds the maximum value in an array of float64s.
//  Returns an error if there are no elements in the array
func Min( vals []float64 ) ( min float64, err error ) {
  if( len( vals ) > 0 ) {
    min = vals[ 0 ];
    err = nil;

    for _, val := range vals {
      if( val < min ) {
        min = val;
      }
    }
  } else {
    err = errors.New( "No values in array" );
  }

  return;
}

//  Finds the minimum value in an array of float64s.
//  Returns an error if there are no elements in the array
func Max( vals []float64 ) ( max float64, err error ) {
  if( len( vals ) > 0 ) {
    max = vals[ 0 ];
    err = nil;

    for _, val := range vals {
      if( val > max ) {
        max = val;
      }
    }
  } else {
    err = errors.New( "No values in array" );
  }

  return;
}
