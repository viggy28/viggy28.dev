---
date: 2021-10-11
description: "Everything that I know of defer statement"
featured_image: ""
tags: ["Go"]
title: "defer statement in go"
---

`defer` is one of the unique keywords in Go.

What does `defer` do?
As the name suggests, it defers something. It defers function calls. A deferred function gets invoked right after the surrounding function returns. 

```Go
func outer() {
    defer inner()
    log.Println("outer got called")
}

func inner() {
    log.Println("inner got called")
}
```

`outer got called` prints first then `inner got called`.

What the main use of `defer` statement?
It helps mainly with actions that need to happen but not immediately. For eg. closing a file that's opened or closing a database connection.

What's the advantage of using `defer`?
- Even when the surrounding function panics, the deferred function call gets invoked
- Makes the code less bug free

Things to remember:

1. The function calls get stacked in Last In First Out (LIFO)

```Go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("1st call")
	}()
	defer func() {
		fmt.Println("2nd call")
	}()
	defer func() {
		fmt.Println("3rd call")
	}()
}
```

The above function outputs:
```bash
viggy28@MacBook-Pro defer % go run defer.go
3rd call
2nd call
1st call
```

2. The (deferred) function’s arguments are evaluated right away, not once the surrounding function returns.

    a. If you would like the function's arguments evaluated at the time of the surround function return then make the argument as a pointer
