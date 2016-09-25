package maths

import "testing"

type testPair struct {
  input []float64
  expectedOutput float64
}

var tests = [] testPair {
  { [] float64 { 1, 2, 3, 4, 5, }, 3 },
  { [] float64 { 3, 6, 9, },       6 },
  { [] float64 { 1, 7, 4 } ,       4 },
  { [] float64 {},                 0 },
};

func TestMean( t *testing.T ) {
  for _, test := range tests {
    v, err := Mean( test.input... );
    _ = err;

    if v != test.expectedOutput {
      t.Error( "Expected ", test.expectedOutput, " got ", v );
    }
  }
}
