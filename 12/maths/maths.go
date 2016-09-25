package maths

import "errors"

//  Finds the mean value of a range of integers
func Mean( vals ...float64 ) ( mean float64, err error ) {
  var total float64 = 0;

  if( len( vals ) > 0 ) {
    for _, x:= range vals {
     total += x;
    }

    mean = total / float64( len( vals ) );
  } else {
    err = errors.New( "Array of values is empty" );
  }

  return;
}
