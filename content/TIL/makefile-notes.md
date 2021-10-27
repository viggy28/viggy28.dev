---
date: 2021-10-26
description: "Types of variables, echo etc"
featured_image: ""
tags: ["make", "linux"]
title: "Notes on makefile"
---

There are two types of variable declaration.

1. recursively expanded variables - `foo` is assigned with `bar` which inturns assigned to `ugh` whose value is `Huh?`

```
foo = $(bar)
bar = $(ugh)
ugh = Huh?

all:;echo $(foo)

$ make
$ Huh?
```

2. simply expanded variables - `foo` is assigned with `bar` however bar is not defined yet. so it outputs to nothing.

```
foo := $(bar)
bar := $(ugh)
ugh := Huh?

all:;echo $(foo)

$ make
$
```

NOTE: To `echo` something it needs to be in a target. Here the target name is `all`.

If I don't want to have target and print something I can use `warning`.

Reference:

https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_6.html

https://www.unix.com/unix-for-dummies-questions-and-answers/104103-how-print-something-make-utility.html