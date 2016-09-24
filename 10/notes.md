#Concurrency

##Goroutines

A goroutine is a function that's capable of running concurrently with other functions

To invoke a function as a goroutine, i.e. asynchronously, use the `go` keyword before it's invocation

e.g.

```
go myFunc( 10 );
```

## Channels

Channels provide a way for goroutines to communicate with each other

Goroutines can synchronize their execution by communicating with each other via a channel

A channel type is declared using the `chan` keyword followed by the type that will be transported via the channel

A variable of type `chan` is assigned calling the `make` function passing a channel type as the first argument

e.g.

```
var c chan string = make( chan string );
```

### Left Arrow Operator

The left arrow operator `<-` is used to send and receive messages to and from channels

e.g.

```
//  send message to channel 'c'

c <- "send this message to the channel c..."

// receive message from channel 'c' and assign it to variable myVar

myVar := <- c 

// receive a message from the channel 'c' and provide it as an argument to Println

fmt.Println( <- c )
```

When you send a message if the channel isn't ready to receive a message, i.e. it's currently processing another message then execution of the goroutine will block to await the channel being ready.

### Channel Direction

You are able to make channels read or write only by specifying 'direction'

e.g.

```
//  write only

c chan<- string

// read only

c <-chan string
```

If a channel doesn't have either of these restrictions it's known as 'bi-directional'

Bi-directional channels can be passed as arguments to functions that have a channel argument that is read/write only but not vice-versa

### Buffered Channels

Where as channels are synchronous by default a buffered channel is asynchronous

You declare a buffered channel by passing a second argument to the `make` function representing the size of the buffer with respect to the type of the channel

e.g.

```
c := make( chan string, 5 );
```

A buffered channel acts like a queue in that it stores messages, up to it's max size, that can be added and removed when needed without blocking.

A buffered channel will block if...
- the channel is full and an attempt is made to add a message to the channel
- the channel is empty and an attempt is made to retreive a messahe from the channel 

### Select

Golang has a `select` control flow structure that is like a `switch` statement but for use with channels

Each `case` statement within a `select` statement sends or receives from the first available channel, blocking if none are available

e.g.

```
select {
  case msg1 := <- c1 :
    fmt.Println( msg1 );
  case msg2 := <- c2 :
    fmt.Println( msg2 );
}
```

# Problems

#How do you specify the direction of a channel type?

By puting the left arrow operator `<-` on the corresponding side of the `chan` keyword

e.g.

```
func myFunc( c chan<- string ) { // write only channel
  //  do stuff...
}

func myFunc( c <-chan string ) { // read only channel
  //  do stuff...
}
```

#Write your own Sleep function using time.After.

- see sleep.go

#What is a buffered channel? How would you create one with a capacity of 20?

- A buffered channel is a type of channel that can hold multiple messages of the declared type.
- Messages can be added and retreived to and from a buffered channel asynchronously
- You can declare a buffered channel with a capacity of 20 like this :

```
myBufferedChannel := make( chan string, 20 );
``` 

- The second argument to the `make` function represents the size of the buffered channel
