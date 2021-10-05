---
date: 2021-10-05T08:45:00UTC-7
description: "A gotcha with respect to deadlocks"
featured_image: ""
tags: ["Go"]
title: "all goroutines are asleep - deadlock!"
---

**program1**
```Go
package main

import (
	"log"
)

func main() {
	done := make(chan struct{})
	go func() {
		log.Println("done")
	}()
	<-done //wait for background goroutine to finish
}
```

**program2**
run `nc -l 8000` on another terminal
```Go
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	done := make(chan struct{})
	go func() {
		log.Println("done")
	}()
	<-done //wait for background goroutine to finish
}
```

Both of those examples have a receive statement from a channel that no one ever sends to. Such a statement must block forever, meaning the goroutine that contains it can never run. If there are no goroutines that can run, Go terminates the program with a deadlock message.

However **program2** wasn't deadlocking. I couldn't figure out why. Later, [Tim](https://github.com/theckman) clarified that

> deadlock detection is disabled in certain situations, including when something in your program imports net.

That explained me why **program2** wasn't deadlocking. 
