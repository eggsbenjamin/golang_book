# Packages

Go was designed, like many other languages, with re-usable components in mind

Packages in Go are synonymous with libraries

Packages are useful for... 
- namespacing related methods
- organising code for re-use
- speeding up compilation as packages remain un-altered 

## Creating Packages

To declare a package you have to...
- use the `package` keyword at the top of the file followed by the name of the package
- create a directory of the same name as your package in `$GOPATH/src`
- put your package file in the directory
- run `$ go install` inside of the newly created directory to compile your package ( if you import the package into another file and do a `$ go run myFile.go` it will automatically compile it for you ) 

__Note__ this is the default way of creating managing packages in Go. There are package managers out there that help to manage packages on a project level rather than a global level

[Go Package Managers](http://jbeckwith.com/2015/05/29/dependency-management-go/)

## Access

There isn't a `private` or a `public` keyword in go.

Instead the convention used is, methods declared with...
- a capital letter as their first letter are public
- a lower case letter as their first letter are private.

## Aliases

You can proivide a custom alias for a package when importing it into another file by using the following syntax :

```
import m "myPackage/package"

//  m now refers to your package rather than it's package mame
```

## Documentation

Go has inbuilt documentation support for packages.

if you run `$ godoc path/to/package/directory MethodName` you'll get documentation for the method printed to stdout

If you run `$ godoc -http=":<port number>"` you'll have a Go documentation web app served locally on the specified port containing information about all packages, including custom packages, installed on your system.

# Problems

## Why do we use packages?

- To group and organise related functionality into reusable components
- To adhere to the rules of encapsulation/abstraction

## What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)

- An identifier that starts with a capital letter is public whereas an identifier that starts with a lower-case letter is private

## What is a package alias? How do you make one?

- A package alias is a custom name for a package other than the default
- you declare a package alias like this `import myAlias "myPackage"` where `myAlias` represents your custom alias and `myPackage` is the name of your package

## We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

- see minMax.go

## How would you document the functions you created in the above question?

- By adding comments above the method declarations 
