package minMax

import "testing"

type testPair struct {
  input [] float64
  expectedOutput float64
}

var minTests = [] testPair {
  { [] float64 { 1, 2, 3,          },  1 },
  { [] float64 { 1, 22, 332,       },  1 },
  { [] float64 { 1, 888, 3,        },  1 },
  { [] float64 { 12, 2333, 3,      },  3 },
  { [] float64 { 1319, 12, 322, -5 }, -5 },
  { [] float64 {},                     0 },
}

func TestMin( t *testing.T ) {
  for _, test := range minTests {
    actualOutput, err := Min( test.input );
    _ = err;

    if( actualOutput != test.expectedOutput ) {
      t.Error( "Expected ", actualOutput, " to be ", test.expectedOutput );
    }
  }
}

var maxTests = [] testPair {
  { [] float64 { 1, 2, 3,          },     3 },
  { [] float64 { 1, 22, 332,       },   332 },
  { [] float64 { 1, 888, 3,        },   888 },
  { [] float64 { 12, 2333, 3,      },  2333 },
  { [] float64 { 1319, 12, 322, -5 },  1319 },
  { [] float64 {},                        0 },
}

func TestMax( t *testing.T ) {
  for _, test := range maxTests {
    actualOutput, err := Max( test.input );
    _ = err;

    if( actualOutput != test.expectedOutput ) {
      t.Error( "Expected ", actualOutput, " to be ", test.expectedOutput );
    }
  }
}
