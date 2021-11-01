---
date: 2021-10-31
description: "Everything that I know of context"
featured_image: ""
tags: ["go"]
title: "Context in go"
---

`context` is a standard package.
`Context` is an interface in `context` package.

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

There are 4 methods in `Context` interface. Most of the times, we don't define our own types which satisfy the interface (which off course we can do), but we mostly use the factory functions provided by the `context` package itself. 

They are `Background()`, `TODO()`, `WithCancel()`, `WithTimeout()` etc. The only difference is WithDeadline sets the termination time point, whileWithTimeout sets the maximum running time. 

`context.TODO()` looks similar to `context.Background()`. `TODO()` is used a filler.

A type which satisfies `Context` interface.

```go
type customContext struct {
	data string
}

func (cc customContext) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), true
}

func (cc customContext) Done() <-chan struct{} {
	var ch chan struct{}
	ch <- struct{}{}
	return ch
}
func (cc customContext) Err() error {
	return errors.New("custom error")
}
func (cc customContext) Value(key interface{}) interface{} {
	return "customvalue"
}
```

1. Deadline

    > Deadline returns the time when work done on behalf of this context
	should be canceled. Deadline returns ok==false when no deadline is
	set. Successive calls to Deadline return the same results.

A simple demonstration of `ctx.Deadline()` method. `ctx.Deadline()` returns a deadline which is a `time.Time` and a `bool` value. Called comma,ok value. A value of false represents there is no deadline set. 


```go
func main() {
	ctx := context.Background()
	fmt.Println(ctx.Deadline())
	fmt.Println(time.Now())
	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()
	fmt.Println(deadlineCtx.Deadline())
	timeoutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	fmt.Println(timeoutCtx.Deadline())
}
```

```bash
0001-01-01 00:00:00 +0000 UTC false
2021-10-31 23:06:48.947153 -0700 PDT m=+0.000413088
2021-10-31 23:06:58.947429 -0700 PDT m=+10.000689674 true
2021-10-31 23:07:08.947486 -0700 PDT m=+20.000746104 true
```

2. Done

    >Done returns a channel that's closed when work done on behalf of this
	context should be canceled. Done may return nil if this context can
	never be canceled. Successive calls to Done return the same value.
	The close of the Done channel may happen asynchronously,
	after the cancel function returns.

	>WithCancel arranges for Done to be closed when cancel is called;
	WithDeadline arranges for Done to be closed when the deadline
	expires; WithTimeout arranges for Done to be closed when the timeout
	elapses.

It is used with `select` statement generally but not necessarily.
`Done()` provides a simple way to know whether a cancellable context has been cancelled or not. calling `Done()` on background or TODO context returns nil, which means the context will never be closed. 

```go
	go func() {
		<-deadlineCtx.Done()
		fmt.Println("deadline context cancelled")
	}()
	go func() {
		<-ctx.Done()
		fmt.Println("background context cancelled")
	}()
	go func() {
		<-timeoutCtx.Done()
		fmt.Println("timeout context cancelled")
	}()
	<-ch
```

3. Err

    > 	If Done is not yet closed, Err returns nil.
	If Done is closed, Err returns a non-nil error explaining why:
	Canceled if the context was cancele d
	or DeadlineExceeded if the context's deadline passed.

    It provides the reason for the context cancellation.

4. Value

    > Value returns the value associated with this context for key, or nil
	 if no value is associated with key. Successive calls to Value with
	 the same key returns the same result.

Contexts are also used for sharing information between functions. Though, generally it is not advisable to pass info through context rather it should be done through function parameters.

Related notes:

1. https://viggy28.dev/gopostgres/go-context/