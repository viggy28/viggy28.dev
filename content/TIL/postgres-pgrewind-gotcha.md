---
date: 2022-09-06
description: "Postgres unable to find replication slot file"
featured_image: ""
tags: ["postgres"]
title: "Postgres pg_rewind gotcha"
---

I use pg_rewind for rewind a Postgres cluster which is ahead of the primary. Yesterday, after the rewind Postgres didn't start. It failed with the below error. This is using 9.6 version pg_rewind client.

```
Sep  2 20:06:14.371503 productiondb[1]: [1-1] time=2022-09-02 20:06:14.338 GMT,pid=78197,user=admin,db=postgres,client=[local],appname=[unknown],vid=,xid=0 FATAL:  the database system is starting up
Sep  2 20:06:14.371611 productiondb[1]: [1-1] time=2022-09-02 20:06:14.370 GMT,pid=78184,user=,db=,client=,appname=,vid=,xid=0 LOG:  entering standby mode
Sep  2 20:06:14.371991 productiondb[1]: [1-1] time=2022-09-02 20:06:14.371 GMT,pid=78184,user=,db=,client=,appname=,vid=,xid=0 PANIC:  could not open file "pg_replslot/ae0750b1/state": No such file or directory
Sep  2 20:06:14.371991 productiondb[1]: [1-1] time=2022-09-02 20:06:14.371 GMT,pid=78182,user=,db=,client=,appname=,vid=,xid=0 LOG:  startup process (PID 78184) was terminated by signal 6: Aborted
```

Of course, the `/state` file didn't exist inside the replication slot sub-directory `ae0750b1`.

Next, I digged the `pg_rewind` log and found an interesting message.

```
Sep 02 19:58:30 productiondb[21494]: received chunk for file "pg_replslot/588e83e5/state", offset 0, size 176
Sep 02 19:58:30 productiondb[21494]: received null value for chunk for file "pg_replslot/ae0750b1/state", file has been deleted
Sep 02 19:58:30 productiondb[21494]: received chunk for file "pg_stat_tmp/db_0.stat", offset 0, size 2371
```

During pg_rewind, replication slot was dropped by another tool. However, it doesn't happen all the time. pg_rewind first copies the file then it copies the content (in chunks). If the replication slot is dropped between the time the file was copied (created) and it's contents are copied then this error `received null value for chunk for file "pg_replslot/ae0750b1/state", file has been deleted occurs.`

Michael Paquier [mentioned](https://www.postgresql.org/message-id/20180124023312.GD1355%40paquier.xyz) a workaround is to remove the slot directory i.e. `ae0750b1` and it's contents (i.e. `state` file) after pg_rewind and before starting postgres.

Of course, Micahel also [committed](https://www.postgresql.org/message-id/20180205071022.GA17337@paquier.xyz) on v11 to skip not only replication slots but also other such files which are not needed by the standby.
