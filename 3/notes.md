#Types

##How are integers stored on a computer?

- As a series of bits representing the corresponding number in binary form
- e.g. 1 = 1, 2 = 10, 3 = 11, 4 = 100, 5 = 101, 6 = 110, 7 = 111, 8 = 1000 etc...

##We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. 
##Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) 
##and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99)

- 11111111 = ( ( 2^8 ) - 1 ) = 255

##Although overpowered for the task you can use Go as a calculator. Write a program that computes 32132 × 42452 and 
##prints it to the terminal. (Use the * operator for multiplication)

- see calculator.go

##What is a string? How do you find its length?

- A string, in the context of GO, is a type i.e. a data structure that holds a series of bytes each representing a character.
- A string will have a length property that corresponds to the number of bytes contained within it.

##What's the value of the expression (true && false) || (false && true) || !(false && false)?

- ( true && false ) || ( false && true ) || !( false && false ) === true 
- This is due to the bracketed part of the expression that evaluates to true 
- It negates, using the ! operator, the statement false && false which evaluates to false.
