---
date: 2019-04-06T18:00:00-00:00
description: "Postgres Active-Active Replication"
featured_image: "/images/postgres-logo.png"
tags: ["postgres", "docker", "haproxy"]
title: "Part 2/2: How to set up HAProxy for an active-active postgres databases"
---
### Step 1: Setting up HAProxy 
I hope you gone through [Part1] (<https://viggy28.dev/article/postgres-active-active-replication/>) of this series. Perhaps, one thing you might have noticed is that I've to connect to the specific master database. In our case, since both the databases are running on docker, only the localhost port is different. (In a production environment, most likely you going to run the databases on a different host). The main reason for active-active replication is high availability. If one of the nodes goes down, you still want your application to behave normally. You don't want to hard-code your DSN on the application or keep checking the health of the database every time before you make a connection. Fair enough. [HAProxy] (<www.haproxy.org>), an open source project solves this particular change.

Basically, you need to connect to a proxy that routes the request to the underlying database hosts. There are different ways you can configure to route the connections. The default is Round-Robin.

The architecture would look something like this:
![alt text](https://gitlab.com/viggy28-websites/viggy28.dev/blob/master/public/images/postgres-haproxy-wb1.jpg "wb1")

You can continue with your docker-compose.yml file. Add the below section

```yml
version: "3"
 
services:
 database0:
   image: jgiannuzzi/postgres-bdr
   restart: always
   ports:
     - 54325:5432
   environment:
     - SERVICE_PORTS=5432
   volumes:
     - /Users/viggy28/tech/docker/volumes/postgres0:/var/lib/postgresql/data
   environment:
     POSTGRES_PASSWORD: "replace with your password"
     TCP_PORTS: "5432"

 database1:
   image: jgiannuzzi/postgres-bdr
   restart: always
   ports:
     - 54326:5432
   environment:
     - SERVICE_PORTS=5432
   volumes:
     - /Users/viggy28/tech/docker/volumes/postgres1:/var/lib/postgresql/data
   environment:
     POSTGRES_PASSWORD: "replace with your password"
     TCP_PORTS: "5432"

 proxy:
   image: dockercloud/haproxy
   links:
     - database0
     - database1
   volumes:
     - /var/run/docker.sock:/var/run/docker.sock
   ports:
     - "15432:5432"
```

Note:

 1. links: links the database0 and database1 services with proxy service
 2. environment: have to export port 5432 where the database is running
 3. ports: in proxy service, port 5432 is forwarded to port 15432 of localhost

### Step 2: Connecting to the HAProxy and verifying its benefit

Verify that all the three services are running

```bash
viggy28@Vigneshs-MacBook-Pro ~ $ docker ps
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                                                NAMES
5a7f75295722        dockercloud/haproxy       "/sbin/tini -- docke…"   About an hour ago   Up About an hour    80/tcp, 443/tcp, 1936/tcp, 0.0.0.0:15432->5432/tcp   postgres0_proxy_1
764dc76bada7        jgiannuzzi/postgres-bdr   "/docker-entrypoint.…"   About an hour ago   Up About an hour    0.0.0.0:54326->5432/tcp                              postgres0_database1_1
89cd08ae8c7f        jgiannuzzi/postgres-bdr   "/docker-entrypoint.…"   About an hour ago   Up About an hour    0.0.0.0:54325->5432/tcp                              postgres0_database0_1
```

Connect to the port 15432 and verify the data

```sql
viggy28@Vigneshs-MacBook-Pro ~ $ psql -h localhost -U postgres -p 15432 -d postgres
Password for user postgres:
psql (11.2, server 9.4.17)
Type "help" for help.

postgres=# \c bdrdemo
psql (11.2, server 9.4.17)
You are now connected to database "bdrdemo" as user "postgres".
bdrdemo=#
bdrdemo=# select * from names;
 user_id |   username   |           email
---------+--------------+---------------------------
       1 | ravichandran | ravikchandran14@gmail.com
(1 row)
```

I am going to stop the container which is running a database on port 54325

```bash
viggy28@Vigneshs-MacBook-Pro haproxy $ docker stop 89cd08ae8c7f
89cd08ae8c7f

viggy28@Vigneshs-MacBook-Pro haproxy $ docker ps
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                                                NAMES
5a7f75295722        dockercloud/haproxy       "/sbin/tini -- docke…"   About an hour ago   Up About an hour    80/tcp, 443/tcp, 1936/tcp, 0.0.0.0:15432->5432/tcp   postgres0_proxy_1
764dc76bada7        jgiannuzzi/postgres-bdr   "/docker-entrypoint.…"   About an hour ago   Up About an hour    0.0.0.0:54326->5432/tcp                              postgres0_database1_1
```

However, I can still able to connect to the database

```bash
viggy28@Vigneshs-MacBook-Pro ~ $ psql -h localhost -U postgres -p 15432 -d postgres
Password for user postgres:
psql (11.2, server 9.4.17)
Type "help" for help.
```

Let me stop the other container which is running a database on 54326

```bash
viggy28@Vigneshs-MacBook-Pro haproxy $ docker stop 764dc76bada7
764dc76bada7
viggy28@Vigneshs-MacBook-Pro haproxy $
viggy28@Vigneshs-MacBook-Pro haproxy $ docker ps
CONTAINER ID        IMAGE                 COMMAND                  CREATED             STATUS              PORTS                                                NAMES
5a7f75295722        dockercloud/haproxy   "/sbin/tini -- docke…"   About an hour ago   Up About an hour    80/tcp, 443/tcp, 1936/tcp, 0.0.0.0:15432->5432/tcp   postgres0_proxy_1
```

Guess what !!?

```bash
viggy28@Vigneshs-MacBook-Pro ~ $ psql -h localhost -U postgres -p 15432 -d postgres
psql: server closed the connection unexpectedly
    This probably means the server terminated abnormally
    before or while processing the request.
```

I hope it makes sense. Basically, all the databases behind the proxy are down.