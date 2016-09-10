#Arrays, Slices and Maps

assignment statements from maps return 2 values:
  - the value associated with they key
  - a boolean indicating if the map contains the key

if statements can include an assignment statement and then check the value of the boolean ( see above ) e.g. 

```
if value, ok := myMap[ "key" ]; ok {
  //  if the map contains the key 'key' then ok will be true
  //  and this block will be evaluated. 

  //  do stuff...
}
```

the above if statement ( ignore the block ) consists of two components delimeted by ';'
  - an assignment statement
  - a check on the ok boolean value returned from the assignment statement

#Problems

##How do you access the 4th element of an array or slice?

using the following syntax: `myArray[ 3 ]`
  - arrays are zero based so the integer 3 is the index of the 4th element

##What is the length of a slice created using: make([]int, 3, 9)?

3, the first 3 elements of the underlying array which has a length of 9

##Given the array:

```
x := [6]string{"a","b","c","d","e","f"}
```

##what would x[2:5] give you?

a slice containing the following elements:

```
[ "c", "d", "e" ]
```

the expression uses ( [startIndex:endIndex] ) returns a slice of the array from, and including the 'startIndex' up to, but not including the 'endIndex'.

##Write a program that finds the smallest number in this list:

```
x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
```

see smallestNumber.go
