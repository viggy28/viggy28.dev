---
date: 2021-11-04
description: "Functional options in Go"
featured_image: ""
tags: ["go"]
title: "functional options"
---

I saw this pattern of a function taking optional parameters (variadic arguments) and I never understood
why they do that until today.

There is a bit of a background on why I need to use function pattern.

```go
package main

import "fmt"

type Config struct {
	Host       string
	Port       uint
	Visibility bool
}

type clientOption func(c *Config) error

func newConfig(host string, opts ...clientOption) Config {
	c := Config{
		Host: host,
	}
	for _, option := range opts {
		option(&c)
	}
	return c
}

func withPort(port uint) clientOption {
	return func(c *Config) error {
		c.Port = port
		if c.Port == 0 {
			c.Port = 5432
		}
		return nil
	}
}

func main() {
	c := newConfig("localhost", withPort(100))
	fmt.Println("config: ", c)
}
```