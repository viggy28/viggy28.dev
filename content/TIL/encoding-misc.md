---
date: 2021-10-30
description: "Misc notes on client encoding, character sets"
featured_image: ""
tags: ["Postgres", "Linux"]
title: "encoding miscellaneous"
---

From the wiki,

```
In computing, data storage, and data transmission, character encoding is used to represent a repertoire of characters by some kind of encoding system that assigns a number to each character for digital representation.
```

Okay, they mean converting character to some number for storing and transmitting data.

There are two popular character sets
1. ASCII 
   1. American Standard Code for Information Interchange
   2. Pretty much all the characters and symbols in modern keyboard comes under ASCII 
   3. Total of 128 characters
   4. Decimal value between 65 and 90 represents captial alphabets and 97 to 122 represents small alphabets
2. UTF-8
    1. ASCII doesn't cover characters from other languages. 
    2. UTF-8 is capable of encoding all 1,112,064 valid character code points in Unicode using one to four one-byte (8-bit) code units.
    3. It has backward compatible with ASCII

I run into character set in postgres. There is `client_encoding` and `server_encoding` setting.

PostgreSQL can determine which character set is implied by the LC_CTYPE setting

`psql` If both standard input and standard output are a terminal, then psql sets the client encoding to “auto”, which will detect the appropriate client encoding from the locale settings (LC_CTYPE environment variable on Unix systems). If this doesn't work out as expected, the client encoding can be overridden using the environment variable PGCLIENTENCODING. ref https://www.postgresql.org/docs/current/app-psql.html#AEN88713

There is also `locale` executable.

UTF-32 stores four bytes. It does waste a lot of memory. So UTF-8 is a better one. Go, uses UTF-8