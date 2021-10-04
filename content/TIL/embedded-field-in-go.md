---
date: 2021-10-04T00:22:00UTC-7
description: "Things that I didn't know about"
featured_image: ""
tags: ["Go"]
title: "Embedded field in Go"
---

1. Initializing embedded field: The field name of an anonymous field during a struct instantiation is the name of the struct (Time)

```Go
type Event struct {
    ID int
    time.Time 
    }

event := Event{
    ID:   1234,
    Time: time.Now(),
}
```

2.  If an embedded field type implements an interface, the struct containing the embedded field will also implement this interface


