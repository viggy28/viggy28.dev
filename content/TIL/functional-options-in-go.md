---
date: 2021-11-04
description: "Functional options in Go"
featured_image: ""
tags: ["go"]
title: "functional options"
---

I saw this pattern of a function taking optional parameters (variadic arguments) and I never understood
why they do that until today.

There is a bit of a background on why I need to use functional pattern.

Let's say I don't use functional option. Then, all the arguments need to be passed. Also, difficult to set default values and implement logic for default values.

```go
package main

import (
	"fmt"
	"log"
)

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
		err := option(&c)
		if err != nil {
			log.Printf("error %v", err)
		}
	}
	return c
}

func withPort(port uint) clientOption {
	return func(c *Config) error {
		if port > 65000 {
			return fmt.Errorf("port can't be higher than 65000")
		}
		c.Port = port
		if c.Port == 0 {
			c.Port = 5432
		}
		return nil
	}
}

func main() {
	c := newConfig("localhost", withPort(100000))
	fmt.Println("config: ", c)
}

```