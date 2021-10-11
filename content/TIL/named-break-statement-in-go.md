---
date: 2021-10-10
description: "A simple example"
featured_image: ""
tags: ["Go"]
title: "Named break statement in Go"
---

I came across this piece of code from [The Go Programming Language](https://www.gopl.io/) book.

```Go
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
```

The complete program is available in [Github](https://github.com/adonovan/gopl.io/blob/master/ch8/du3/main.go)

Named `break` statement allows to break perhaps multiple control statements. In this example, `break loop` breaks the `select` and `for` statement.