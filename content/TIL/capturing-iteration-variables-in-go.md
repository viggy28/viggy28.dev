---
date: 2021-10-09
description: "A caveat with respect to using loop variables"
featured_image: ""
tags: ["Go"]
title: "capturing iteration variables in Go"
---

I was going through The Go Programming Language and in chapter 5 there was a section `Caveat: Capturing Iteration Variables`.

The example is like this:

```Go
var rmdirs []func()

for _, dir := range tempDirs() {                    // 1
    os.MkdirAll(dir, 0755)                          // 2
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!       // 3
    })
}
```

1. `dir` is the same variable. i.e. the address of the variable is same
2. `dir`'s value evaluated at this point is different at each iteration
3. **IMPORTANT** `dir` is not evaluated here. Its just being used. It will be evaluated only when the `func()` is called

The following explanation from Ref[1]:

```Go
var rmdirs []func()
tempDirs := []string{"one", "two", "three", "four"}

for _, d := range tempDirs {
    dir := d
    fmt.Printf("dir=%v, *dir=%p\n", dir, &dir)
    rmdirs = append(rmdirs, func() {
        fmt.Printf("dir=%v, *dir=%p\n", dir, &dir)
    })
}

fmt.Println("---")

for _, f := range rmdirs {
    f()
}
```
![explanation](/images/TIL-Go-loop-iteration.png)

Reference:
1. [https://stackoverflow.com/questions/52980172/cannot-understand-5-6-1-caveat-capturing-iteration-variables](https://stackoverflow.com/questions/52980172/cannot-understand-5-6-1-caveat-capturing-iteration-variables)
2.  [https://github.com/golang/go/wiki/CommonMistakes](https://github.com/golang/go/wiki/CommonMistakes)




