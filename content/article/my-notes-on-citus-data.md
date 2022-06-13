---
date: 2022-06-12
description: "Notes from citus data sigmod 2021 white paper"
featured_image: "/images/postgres-logo.png"
tags: ["postgres", "distributed systems"]
title: "Notes from citus data white paper"
---

[Citus: Distributed PostgreSQL for Data-Intensive Applications](https://dl.acm.org/doi/pdf/10.1145/3448016.3457551)

Context: Recently, our team got a request to provide a solution to shard Postgres. One of the solutions that was discussed Citus. I have heard about them and seen their blogs related to Postgres in the past but never used the product. I thought it would be fun to read about its internal working.

Before digging their white paper, taking a step back and asking what is sharding and why do I need to shard?

Certainly, I haven't used the word "shard" in my ![day-to-day](/images/my-notes-on-citus-data-1.png) life. In CS, there two ways one can scale:

    1. Vertical Scaling     - Adding more resources (eg. CPU, Memory, Disk) on the **same** hardware (eg. server, switch)
    2. Horizontal Scaling   - Adding more resources by adding more hardware (eg. server)

Sharding comes from the concept of Horizontal Scaling. I have seen servers with a maximum of 22 TB disk. What if I want to store more than that in a single table/database? Traditional approach is vertical scaling i.e) trying to add more disks on the server, but at some point it will hit the ceiling. Nothing one can do other than growing horizontally. [A good introduction about database sharding from Digital Ocean](https://www.digitalocean.com/community/tutorials/understanding-database-sharding).

What is Citus?

1. It is a Postgres extension to store data, query (which includes transactions) acorss a cluster of PostgreSQL servers.
2. [Open sourced](https://github.com/citusdata/citus)
 
Postgres core itself doesn't come with features for horizontal scaling. [Postgres' wiki on sharding](https://wiki.postgresql.org/wiki/WIP_PostgreSQL_Sharding#:~:text=PostgreSQL%20provides%20a%20number%20of,pushed%20down%20to%20the%20shards) and [Gitlab's experiment using FDW](https://about.gitlab.com/handbook/engineering/development/enablement/database/doc/fdw-sharding.html) are good resources.

Alternate approaches are:

1. Build the database engine from scratch and write a layer to provide over-the-wire SQL compatibility - aka [YugaByte](https://github.com/yugabyte/yugabyte-db), [Cockroachdb](https://github.com/cockroachdb/cockroach) etc.
2. Fork an open source database systems and build new features on top of it - aka [Orioledb](https://github.com/orioledb/orioledb), [Neondatabase](https://github.com/neondatabase/neon)
3. Provide new features through a layer that sits between the application and database, as middleware - aka [ShardingSphere](https://github.com/apache/shardingsphere)

The types of applications that requires distributed postgres is broadly divided into four categories:

1. Multi-tenant/SaaS
    * An application which stores data of multiple tenants in the same database.
    * Data is relatively specific to the tenant
    * Traditional approach (application level sharding) is spinning up individual database/server for each tenant and then mapping that information on the application itself. There is an operation overhead when moving data around, schema changes and doing analytics across tenants.
    * The alternative approach is the database level sharding. Application doesn't need to track which tenant is stored where. Use shared schema with tenant ID columns. The dbms should be capable of routing arbitrarily complex SQL queries of a specific tenant to a specific server. Should provide support for flexible data type (achieved using JSONB) and control over tenant placements to avoid noisy-neighbor problems

![An example of messaging system which stores multiple tenant data](/images/citus-slack.png). AKA slack.

2. Real-time analytics
    * Used for system monitoring, ingesting IoT data, user browsing/behavioral data etc.
    * System should be capable of supporting parallel bulk loading, INSERT..SELECT to create rollup tables.

3. High-Performance CRUD
   * CRUD stands for Create, Read, Update and Delete
   * An example of such system will be an e-commerce website
     * Generally, they are highly concurrent, expects low latency and needs to do joins

4. Data Warehousing
    * Combines data from different sources into a single database system to generate ad-hoc reports
    * Generally don't have low latency, high concurrency requirements

How citus provides solutions for the above use cases is the rest of the paper. That's basically the nitty-gritty of citus.

**Architecture**
All servers in a Citus cluster, runs vanilla PostgreSQL. It uses extensions api to change the behavior. It replicates [custom types](https://www.postgresql.org/docs/current/sql-createtype.html) and functions across all servers. Adds two new tables a) Distributed table b) Reference table

a. PostgreSQL extension APIs
This is perhaps one of the best features of Postgres. One can change the behavior of PostgreSQL by defining hooks (custom logic). AFAIK, Oracle, MySQL does not have extensions. 

Citus the following hooks 

1. User-defined functions
Callable from SQL inside a transaction usually to manipulate Citus metadata

2. Planner and executor hooks
Citus checks whether the query involves a Citus table, if so intercepts it and creates a plan that contains a CustomScan node
    a. CustomScan is an execution node in the query plan. It calls the Citus query executor which returns results then that will be returned to Postgres query executor.


Open Questions:
1. What's the difference (in features) between open source Citus and paid?