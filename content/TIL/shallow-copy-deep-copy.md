---
date: 2021-10-02T10:00:08-05:00
description: "Gotchas that I ran into"
featured_image: ""
tags: ["Go"]
title: "Understanding shallow copy vs deep copy in Go"
---

In Go, copying generally does a copy by value. So modifying the destination, doesn't modify the source. However, that holds good only if you are copying value types not reference types (i.e. pointer, slice, map etc.)

```Go
package main

import (
	"fmt"
)

func main() {

	type Cat struct {
		age     int
		name    string
		Friends []string //reference type
	}
	wilson := Cat{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}
	// shallow copy
	nikita := wilson
	// modifies both nikita and wilson
	nikita.Friends[0] = "cruise"
	fmt.Println("nikita:", wilson)
	fmt.Println("wilson:", wilson)
	// Here I am modifying the slice too, but whats stored in nikita.Friends is the address of new slice which append returns not wilson.Friends
	nikita.Friends = append(wilson.Friends, "jerry")
	fmt.Println("nikita:", nikita)
	fmt.Println("wilson:", wilson)
	// If the above doesn't makes sense this should. Created a new slice data
	data := []string{"popeye"}
	// shallow copy data to nikita Friends
	nikita.Friends = data
	// changing nikita Friends should change data too
	nikita.Friends[0] = "pluto"
	fmt.Println("nikita:", nikita)
	fmt.Println("wilson:", wilson)
	fmt.Println("data:", data)
}

```

This code explains one of the gotchas that I ran into.
`nikita.Friends = append(wilson.Friends, "jerry")` I was expecting `wilson.Friends` will also have `jerry` but since i did shallow copy `nikita := wilson`. However, that's not the case.
I realized `nikita.Friends = append(wilson.Friends, "jerry")` actually creates a new slice and stores that's reference. The `data` should clarify that further.

Another simpler example to understand the side effects of shallow copy in reference types

```Go
package main

import (
	"fmt"
)

func main() {

	type Cat struct {
		age     *int
		name    string
		Friends []string //reference type
	}
	catAge := 4
	wilson := Cat{&catAge, "Wilson", []string{"Tom", "Tabata", "Willie"}}
	// shallow copy
	nikita := wilson
	*nikita.age = 3
	fmt.Println("nikita:", *nikita.age)
	fmt.Println("wilson:", *wilson.age)
	type Cats struct {
		age     int
		name    string
		Friends []string //reference type
	}
	wilson1 := Cats{4, "Wilson", []string{"Tom", "Tabata", "Willie"}}
	// shallow copy
	nikita1 := wilson1
	nikita1.age = 3
	fmt.Println("nikita1:", nikita1.age)
	fmt.Println("wilson1:", wilson1.age)
}
```
