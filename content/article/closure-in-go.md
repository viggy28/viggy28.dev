---
date: 2021-10-30
description: "Everything that I know of closure"
featured_image: ""
tags: ["Go", "closure"]
title: "Closure in go"
---

closure is a function inside another function where you can reference variable defined in the outer function.

I have seen them useful in dealing with packages where a function expects an `type func()` as an argument where the `func()` doesn't take any argument.

example from package github.com/cenkalti/backoff

signature of backoff.Retry is as below

```go
backoff.Retry(o backoff.Operation, b backoff.Backoff)
```

where backoff.Operation is of type `func() error`

```go
type Operation func() error
```

```go
func main() {
    func retryWrapper(imp string) func() error{
        return func() {
            funcThatCouldFail(imp)
            retun nil
        }
    }
    err := backoff.Retry(retryWrapper("hello world"), backoff.NewExponentialBackOff())
}
```

we can't directly call `funcThatCouldFail(imp)` in `backoff.Retry()`. So, we create a wrapper which takes an argument and that argument is being used inside of the func that being returned.