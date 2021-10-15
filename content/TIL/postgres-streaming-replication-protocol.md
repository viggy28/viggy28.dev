---
date: 2021-10-13
description: ""
featured_image: ""
tags: ["Postgres"]
title: "Postgres streaming replication protocol"
---

I have know about the postgres wire protocol, but first I ran into streaming replication protocol.

I was looking at this code in [stolon](https://github.com/sorintlab/stolon/blob/adf3437d7e02184024dc110522648e915451cb52/internal/postgresql/utils.go#L229)

```Go
	replConnParams["replication"] = "1"
	db, err := sql.Open("postgres", replConnParams.ConnString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := query(ctx, db, "IDENTIFY_SYSTEM")
	if err != nil {
		return nil, err
	}
```

I was wondering what is this `IDENTIFY_SYSTEM`. A G-search pointed me to [this](https://www.postgresql.org/docs/9.5/protocol-replication.html)

```bash
visi@viggy28-MacBook-Pro ~ % psql "dbname=postgres user=postgres replication=database password=replaceme" -c "IDENTIFY_SYSTEM;"
      systemid       | timeline |  xlogpos  |  dbname
---------------------+----------+-----------+----------
 6852440194718025990 |        1 | 0/1542DB8 | postgres

visi@visis-MacBook-Pro ~ % psql "dbname=postgres user=postgres replication=database password=replaceme" -c "TIMELINE_HISTORY 1;"
ERROR:  could not open file "pg_xlog/00000001.history": No such file or directory
(1 row)
```

I learnt you can directly postgres using other than SQL by sending messages to the backend (server) through the wire protocol !!

The systemid is same for all the instances (servers) in a cluster (primary, replica)
