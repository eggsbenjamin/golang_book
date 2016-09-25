# Testing

In Go tests are identified by starting with the word Test and taking one paramter of type `*testing.T`

To run tests use the `$ go test` command

This will...
- run all tests in the current directory
- outut results to stdout

When writing tests make sure to...
- declare the same package name as that of the file that contains the method that you are testing
- import the `testing` package

## Fixtures

A common way to set up tests is to declare a struct containing...
- the input for the method
- the expected output

This way you can loop through the fixtures and assert against their actual output

# Problems

## Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about a problem then you may at first realize. For example, with our Mean function what happens if you pass in an empty list ([]int{})? How could we modify the function to return 0 in this case?

- You could just check the length of the input array and explicitly return 0 if it's less than 1
- You could have an error as a second return type from the method and return an error if the length is less than 1. In this case the default value for the flat64 is returned ( 0 ) along with a meaningful error message 

## Write a series of tests for the Min and Max functions you wrote in the previous chapter.

- see minMax/minMax_test.go
