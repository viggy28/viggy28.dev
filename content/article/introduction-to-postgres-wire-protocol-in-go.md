---
date: 2023-04-21T18:00:08-04:00
description: "A gentle introduction to postgres wire protocol in Go"
featured_image: ""
tags: ["postgres", "go"]
title: "pglite"
---

I kept hearing about the term wire protocol especially Postgres wire protocol in the recent days (Looking at you cockroachdb, yugabytedb - in a good way) but never really quite understood it. Decided to implement something simple in Go to understand it better. As always, if you find anything wrong or I misunderstood please correct me.

Debunking these complex terms.

"wire"      - something over network generally (but PG also supports over domain sockets)

"protocol"  - contract or an agreement between frontend and backend 

Refer https://www.postgresql.org/docs/current/protocol.html and https://www.postgresql.org/docs/current/protocol-flow.html. If you are like me who has seen those links multiple times but didn't know how to implement something using that, then this is for you :) Also, if you just want to see the code it is [here](https://github.com/viggy28/pglite). 

## step 1: Create a TCP server on port 5432
```go
	listenAddr := net.TCPAddr{
		Port: 5432,
	}
	listener, err := net.ListenTCP("tcp", &listenAddr)
	if err != nil {
		log.Fatalf("error while trying to start the server %v", err)
	}
```

## step 2: Continously accept concurrent connections
```go
	for {
		conn, err := listener.Accept()
		if err != nil {
			ErrorLogger.Printf("error accepting connection from %s %s %v", conn.RemoteAddr().Network(), conn.RemoteAddr().String(), err)
			continue
		}
		InfoLogger.Println("received a new connection from ", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
		// handle the client connection in a separate go routine so we don't block subsequent client connections
		go handler(db, conn)
	}
```

## step 3: Handle connections from Postgres client (i.e. psql etc.)

The first step is "startup" phase. During which the byte order is as below

```
    //
	// byte ordering
	// message-length		message
	// [][][][] 			[][][][][][]....
	// message-length count includes its own byte value. For eg. if the content of message-length is [0 0 0 84]
	// then 80 is the number of message byte which contains the message (ie. subtract 4 since that's the number of bytes it takes to store an integer)
	// 
```

```go
    readBuf := make([]byte, 4)
	_, err := c.Read(readBuf)
	if err != nil {
		return fmt.Errorf("unable to read data from buffer %v", err)
	}
	// convert (aka decode) data stored in bytes to integer
	msgSize := binary.BigEndian.Uint32(readBuf) - 4
	InfoLogger.Printf("startup message size including message-length: %d", int(binary.BigEndian.Uint32(readBuf)))
	if msgSize < 4 || msgSize > 10000 {
		return fmt.Errorf("invalid length of startup message size: %d", msgSize)
	}
	msgBuf := make([]byte, msgSize)
	_, err = c.Read(msgBuf)
	if err != nil {
		return fmt.Errorf("unable to read data from message buffer %v", err)
	}
    message := binary.BigEndian.Uint32(msgBuf)
```

If you are here, that means you have successfully parsed a startup message from Postgres client. Next, need to figure how to handle that. Showing only the authentication part.

```Go
	// message format looks like below for AuthenticationOk(B)
	// ['R'] [][][][8] [][][][0]
	c.Write([]byte("R"))
	messageLen := 8
	lenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lenByte, uint32(messageLen))
	if _, err := c.Write(lenByte); err != nil {
		return fmt.Errorf("error writing lenByte %v in AuthenticationOk", err)
	}
	successfulAuth := 0
	authByte := make([]byte, 4)
	binary.BigEndian.PutUint32(authByte, uint32(successfulAuth))
	if _, err := c.Write(authByte); err != nil {
		return fmt.Errorf("unable to write auth byte")
	}
```

Receiving data from client and read it. Based on that creating bytes and sending data back is pretty much the wire protocol is all about. If you find it cumbersome, you can use libraries which abstracts some of those.

## Step 4: Handle client queries

The byte order after startup phase is like below:

```
	// byte ordering
	// message-type		message-length		message
	// []  				[][][][] 			[][][][][][]....
	// Note: message-length count includes itself
	// Note: The very first message sent by the client after the startup message has an message-type byte.
```

```Go
	readBuf := make([]byte, 5)
	// Read() blocks until it can read from conn
	_, err := c.Read(readBuf)
	if err != nil {
		ErrorLogger.Printf("unable to read data from buffer %v \n", err)
		return
	}
    msgType := readBuf[0]
	if msgType == byte(MessageTypeTerminate) {
		return
	}
```

Rest of the message decoding of a Query phase is similar to startup phase. Once you have the query, then you can handle it. I decided to use sqlite as storage engine since it's an embeddable database. At this point the project started sounding similar to [postlite](https://github.com/benbjohnson/postlite). Anyway, I continued since I liked the rabbit hole that I was digging :)

## Step 5: handling SELECT statements

Creating and filling bytes becoming cumbersome and lengthy. Thanks to Jackc who has a library in Go for the Postgres wire protocol.
https://github.com/jackc/pgproto3

To respond to a SELECT query, we need two things. 

1. Field description (i.e. column names and its type etc)
2. Data Row (i.e. actual data returned by SELECT)

```go
	fd = append(fd, pgproto3.FieldDescription{
		Name:                 fieldName,
		TableOID:             0,
		TableAttributeNumber: 0,
		// mapping sqlite types to Postgres' type OID and size
		DataTypeOID:  uint32(dataTypeLookup[fieldType].DataTypeOID),
		DataTypeSize: int16(dataTypeLookup[fieldType].DataTypeSize),
		TypeModifier: -1,
		Format:       0,
	})
```

```go
	for _, row := range resultSet {
	buf = (&pgproto3.DataRow{
			Values: row,
			}).Encode(buf)
	}
```

`row` represents row value from the SELECT query. Finally, write a response back to the client.

```go
	buf = (&pgproto3.ReadyForQuery{TxStatus: 'I'}).Encode(buf)
	if _, err := c.Write(buf); err != nil {
		ErrorLogger.Printf("unable to write response to client %v", err)
	}
```

Similarly, I implemented other statements such as `INSERT`, `UPDATE`, `DELETE`. Complete code is available https://github.com/viggy28/pglite

References:
1. https://gavinray97.github.io/blog/postgres-wire-protocol-jdk-21
2. https://github.com/jackc/pgproto3/tree/master/example/pgfortune
3. https://15721.courses.cs.cmu.edu/spring2023/slides/15-networking.pdf
