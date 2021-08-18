---
date: 2021-08-17T20:20:08-05:00
description: "What is USR2 and where it is used"
featured_image: ""
tags: ["system"]
title: "USR2 kill signal in Linux"
---

I have used many signals with `kill` like `SIGHUP`, `SIGKILL` etc. Today I came across `USR2`. 

Did a Google search and found the [manual page from GNU](https://www.gnu.org/software/libc/manual/html_node/Miscellaneous-Signals.html). According to it,

```
These signals are used for various other purposes. In general, they will not affect your program unless it explicitly uses them for something.
```

Another one from the [BSD mailing list](https://lists.freebsd.org/pipermail/freebsd-questions/2007-August/156889.html)

```
USR2 is a "user defined signal" (from "man signal")

It doesn't "mean" anything by definition.  Each application is free to define its meaning as it sees fit.  It's there specifically so that applications can use signals for special purposes without reusing the defined signals.
```
