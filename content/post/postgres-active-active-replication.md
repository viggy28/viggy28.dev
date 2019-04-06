---
date: 2019-04-05T19:00:00-00:00
description: "Postgres Active-Active Replication"
featured_image: "/images/firebase-logo-wide.png"
tags: ["postgres", "docker", "bdr"]
title: "How to set up active-active replication in postgres using BDR"
---

Postgres doesn't support active-active replication natively. As of this writing, we have to rely on 3rd party tools. I decided to go with [BDR] (<https://www.2ndquadrant.com/en/resources/postgres-bdr-2ndquadrant/>).

I didn't want to spin up multiple VMs. So, the obvious choice is docker. Make sure you have docker on mac & docker compose.

Step1: Running 2 postgres instances

Thanks to [jgiannuzzi] (<https://github.com/jgiannuzzi>). He created a docker [image] (<https://hub.docker.com/r/jgiannuzzi/postgres-bdr>) with postgres and bdr.

docker-compose.yml file content

```
version: "3"

services:
 database0:
   image: jgiannuzzi/postgres-bdr
   restart: always
   ports:
     - 54325:5432
   volumes:
     - /Users/viggy28/tech/docker/volumes/postgres0:/var/lib/postgresql/data
   environment:
     POSTGRES_PASSWORD: <replace with your password>
 database1:
   image: jgiannuzzi/postgres-bdr
   restart: always
   ports:
     - 54326:5432
   volumes:
     - /Users/viggy28/tech/docker/volumes/postgres1:/var/lib/postgresql/data
   environment:
     POSTGRES_PASSWORD: <replace with your password>
```

Things to note:
    ports - I am using my local machine's port 54325 and 54326.
    volumes - postgresql/data is mounted on postgres0 and postgres1 directory.

execute docker-compose up to create the two containers

You can verify that using docker ps

viggy28@Vigneshs-MacBook-Pro ~ $ docker ps
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                     NAMES
0213199e6d0b        jgiannuzzi/postgres-bdr   "/docker-entrypoint.…"   14 minutes ago      Up 9 seconds        0.0.0.0:54326->5432/tcp   postgres0_database1_1
beca9adb4b65        jgiannuzzi/postgres-bdr   "/docker-entrypoint.…"   14 minutes ago      Up 10 seconds       0.0.0.0:54325->5432/tcp   postgres0_database0_1

Check you can connect to the database running on both the containers using

```
psql -h localhost -U postgres -p 54325 -d postgres
psql -h localhost -U postgres -p 54326 -d postgres
```

Thats basically running postgres using docker container.

Step2: Setting up active-active replication using BDR
Connect to database0 (running on port 54325)

```
create database bdrdemo;

\c bdrdemo

CREATE EXTENSION IF NOT EXISTS btree_gist;
CREATE EXTENSION IF NOT EXISTS bdr;
SELECT bdr.bdr_group_create(
local_node_name := 'postgres0_database0_1',
node_external_dsn := 'host=postgres0_database0_1 port=5432 dbname=bdrdemo password=replace with your password'
);

select * from bdr.bdr_nodes;
```

Connect to database1 (running on port 54326)

```
create database bdrdemo;

\c bdrdemo

CREATE EXTENSION IF NOT EXISTS btree_gist;
CREATE EXTENSION IF NOT EXISTS bdr;

SELECT bdr.bdr_group_join(
    local_node_name := 'postgres0_database1_1',
    node_external_dsn := 'host=postgres0_database1_1 port=5432 dbname=bdrdemo password=replace with your password',
    join_using_dsn := 'host=postgres0_database0_1 port=5432 dbname=bdrdemo password=replace with your password'
);

select * from bdr.bdr_nodes;
```

Things to note:
    local_node_name or host name is the container name
    Don't forget to replace your password
    If you face issue with connectivity, make sure one container can communicate with the other one (ping command should be good)
    
```
    $ docker exec -it postgres0_database0_1 /bin/bash
    root@beca9adb4b65:/# ping postgres0_database1_1
64 bytes from 172.20.0.3: icmp_seq=0 ttl=64 time=0.099 ms
64 bytes from 172.20.0.3: icmp_seq=1 ttl=64 time=0.154 ms
64 bytes from 172.20.0.3: icmp_seq=2 ttl=64 time=0.113 ms
^C--- postgres0_database1_1 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
```

Step3: Verifying that data is getting replicated

On database0 (running on port 54325)

```
CREATE TABLE names(
 user_id serial PRIMARY KEY,
 username VARCHAR (50) UNIQUE NOT NULL,
 email VARCHAR (355) UNIQUE NOT NULL
);

insert into names (user_id, username, email) values (1, 'ravichandran', 'ravikchandran14@gmail.com');
INSERT 0 1

```

On database1 (running on port 54326)

```
bdrdemo=# select * from names;
 user_id |   username   |           email           
---------+--------------+---------------------------
       1 | ravichandran | ravikchandran14@gmail.com
(1 row)
```

voila. A simple way to set up active-active replication in postgres.

