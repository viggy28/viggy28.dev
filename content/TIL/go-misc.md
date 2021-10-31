---
date: 2021-10-31
description: "Misc notes on Go"
featured_image: ""
tags: ["Go"]
title: "Misc notes on Go"
---

1. go run command
    `go run` command is a quick way to run go program. However, what it does behind the screen is it compiles the program and runs an executable and removes it once the program is stopped.

2. byte is a built-in alias of uint8. 
   rune is a built-in alias of int32. 

3. The semicolon insertion rule
    In Go, I thought there is no need for semicolon at the end of the statement. It is true (as a developer I don't need to) however lexer (which I don't know what it is; sounds similar to compiler, interpretor) does that behind the screen. Refer https://golang.org/doc/effective_go#semicolons

4. $GOPATH/go.mod exists
   Ran into this when building a go program inside a Go docker image. Figured the reason is my program has modules (go.mod and go.sum file) which copied into container directory which is same as the GOPATH.
   solution: copy the program to a path other than GOPATH