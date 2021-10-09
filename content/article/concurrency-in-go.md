
---
date: 2021-10-04T10:00:00-07:00
description: "includes channels, mutexes"
featured_image: ""
tags: ["Go"]
title: "Notes on concurreny in Golang"
---

### Concurrency

channels are like typed pipes, Where you can send something from one side and receive from other side.
using channel direction operator <-. Depends on the direction of the variables around the operator, its either sending or receiving.
x  := <-c // receive from c
c <- sum // send sum to c
fmt.Println("channel receive happening", <-c)
channels are reference type. similar to maps. So when we pass channel as an argument to a function or copy we are passing the reference.

When the main routine (i.e. main program) returns or when you exit the program then all other go routines will be abruptly terminated.
This example https://github.com/adonovan/gopl.io/blob/master/ch8/spinner/main.go where once the Fibonacci value is calculated it abruptly stops spinner()


> "By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables"
This is something I can't understand
Update: In unbuffered channel if you send something and don't receive from another routine than all operations after the send on that routine will be blocked.
Similarly on a routine if you try to receive from a channel when value is not sent then it blocks



One thing which I tested, if you try to receive more than what you sent in a channel, then it fails with
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
```
Buffered chan:
A channel with defined size is buffered channel. It blocks at the sender when the channel is full and at the receiver when the channel is empty. A good usage of buffered channel from GOPL book

```Go
    func mirroredQuery() string {
        responses := make(chan string, 2)
        go func(){
            responses <- request("asia.gopl.io")
        }()
        go func(){
            responses <- request("europe.gopl.io")
        }()
        out := <- responses
        return out
    }

    func request(hostname string) (response string) {
        // To fill in
    }
```

Whichever `responses` come first that will be returned.


https://replit.com/@viggy28/concurrency#main.go
https://replit.com/@viggy28/concurrency-1#main.go

Also, didn't understand the select statement example https://tour.golang.org/concurrency/5

Channels are good for communication between goroutines. However, we don't need to communicate all the time, rather we just need to make sure one goroutine is accessing a variable at a time. That's where mutex helps.