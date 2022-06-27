---
date: 2022-06-12
description: "Notes from Citus Data sigmod 2021 white paper"
featured_image: "/images/postgres-logo.png"
tags: ["postgres", "distributed systems"]
title: "Citus Data - How it enables distributed postgres"
---

Citus: Distributed PostgreSQL for Data-Intensive Applications paper can be downloaded [here](https://dl.acm.org/doi/pdf/10.1145/3448016.3457551).

> Recently, our team got a request to provide a solution to shard Postgres. One of the solutions that we discussed was [Citus](https://www.citusdata.com/). I have heard about the product and seen their blogs related to Postgres in the past but never used it. I thought it would be fun to read about its internal workings.

If you find something wrong on the notes, please send a [pull request](https://github.com/viggy28/viggy28.dev/tree/master/content/article). Before digging their white paper, let's take a step back and ask what is sharding and why do we need sharding?

Certainly, I haven't used the word "shard" in my day-to-day life. ![shard-meaning](/images/my-notes-on-citus-data-1.png).

There two ways one can scale their systems:

1. Vertical Scaling - Acquiring more resources (eg. CPU, Memory, Disk) on the **same** hardware
2. Horizontal Scaling - Acquiring more resources by adding additional hardware

Sharding comes from the concept of Horizontal Scaling. Say, the maximum disk space of servers that you have is 11 TB. What if we want to store more than that in a single table/database? Traditional approach is vertical scaling i.e) trying to add more disks on the server, but at some point it will hit the ceiling. Nothing one can do other than growing horizontally. [A good introduction about database sharding from Digital Ocean](https://www.digitalocean.com/community/tutorials/understanding-database-sharding).

### What is Citus?

1. It is a PostgreSQL extension to store data, query (which includes transactions) acorss a cluster of PostgreSQL servers
2. [Open sourced](https://github.com/citusdata/citus) [1]

Postgres core itself doesn't come with features for horizontal scaling. [Postgres' wiki on sharding](https://wiki.postgresql.org/wiki/WIP_PostgreSQL_Sharding#:~:text=PostgreSQL%20provides%20a%20number%20of,pushed%20down%20to%20the%20shards) and [Gitlab's experiment using FDW](https://about.gitlab.com/handbook/engineering/development/enablement/database/doc/fdw-sharding.html) are good resources.

Alternate approaches are:

1. Build the database engine from scratch and write a layer to provide over-the-wire SQL compatibility - [YugaByte](https://github.com/yugabyte/yugabyte-db), [Cockroachdb](https://github.com/cockroachdb/cockroach) etc.
2. Fork an open source database systems and build new features on top of it - [Orioledb](https://github.com/orioledb/orioledb), [Neondatabase](https://github.com/neondatabase/neon)
3. Provide new features through a layer that sits between the application and database, as middleware - [ShardingSphere](https://github.com/apache/shardingsphere)

![citus-intro-tweet](/images/citus-tweet-1.png).

I couldn't find a lot of options for horizontal scaling Postgres. [Looks like many agrees](https://twitter.com/viggy28/status/1536157371465990144). MySQL has [Vitess](https://github.com/vitessio/vitess)

The types of applications that requires distributed postgres is broadly divided into four categories:

1. Multi-tenant/SaaS
    * An application which stores data of multiple tenants in the same database.
    * Data is relatively specific to the tenant
    * Traditional approach (application level sharding) is spinning up individual database/server for each tenant and then mapping that information on the application itself. There is an operation overhead when moving data around, performing schema changes and analytics across tenants.
    * The alternative approach is the database level sharding. Application doesn't need to track which tenant is stored in which server. Use a shared schema with tenant ID columns. The dbms should be capable of routing arbitrarily complex SQL queries of a specific tenant to a specific server. Should provide support for flexible data type (achieved using JSONB) and control over tenant placements to avoid noisy-neighbor problems

An example of a messaging system which stores multiple tenant data. AKA slack. ![An example multi tenant messaging app](/images/citus-slack.png).

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

How citus provides solutions for the above use cases is the rest of the paper.

## Architecture
All servers in a Citus cluster, run PostgreSQL with Citus extension. It has two components -- Coordinator and Worker. Typically set up will have 1 Coordinator and 0 or more workers. Coordinator can also be scaled if throughput becomes the bottleneck. If there is no worker then the Coordinator will take that role.

It uses extensions api to change the behavior. It replicates [custom types](https://www.postgresql.org/docs/current/sql-createtype.html) and functions across all servers.

PostgreSQL extension APIs
This is perhaps one of the best features of Postgres. One can change the behavior of PostgreSQL by defining hooks (custom logic). AFAIK, Oracle, MySQL does not have extensions. Citus uses the following hooks

1. User-defined functions
Callable from SQL inside a transaction usually to manipulate Citus metadata

2. Planner and executor hooks
Citus checks whether the query involves a Citus table, if so intercepts it and creates a plan that contains a **CustomScan** node
    a. CustomScan is an execution node in the query plan. It calls the Citus distributed query executor which returns results then that will be returned to Postgres query executor.

3. Transaction callbacks, utility hook and background workers are other hooks used by Citus.

![citus-architecture](/images/citus-architecture.png)

Citus has two types of tables:

**1. Distributed table**
    They are hash-partitioned along a distribution column into multiple logical shards with each shard containing a contiguous range of hash values. From the above diagram, *items* and *users* table are distributed tables with distributed column of *user_id*. One worker node can contain multiple logical shards, so that they can be rebalanced.

**2. Reference table**
    These are replicated to all nodes including coordinator. Joins between distributed tables and reference tables used the local replica of the reference table. From above, *Categories* is a reference table.

I see a semblance of [Star schema](https://en.wikipedia.org/wiki/Star_schema) (fact and dimension tables).

### Co-location:
Citus can make sure that the same range hash values are always on the same worker node among distributed tables. From above, users_4, items_4 (both have hash value of 4) will reside on the same worker node [2]. Main benefit is joins and foreign keys are implemented within a worker node.

### Data rebalancing:
AKA shard rebalancing.
By default, the rebalancer moves shards until it reaches an even number of shards across worker nodes. Also, one can rebalance based on data size or using custom definitions by cost, capacity and constraint function. They do that using PostgreSQL logical replication [3].

### Distributed Query planner and executor:
Fast path planner handles queries on a single table with a single distribution column value.
Router planner handles complex queries that can be scoped to one set of co-located shards.
Logical planner handles queries across shards by constructing a multi-relational algebra tree.

Executor runs in parallel by opening multiple connections per shard instead of using PostgreSQL [parallel query](https://www.postgresql.org/docs/current/parallel-query.html) capability. They found that it is more versatile and performant however with the downside of opening multiple connections. PostgreSQL connections are expensive. They avoid using "slow start" - technique to open a new connection for ever 10ms. Also, paper specified that Crunchy is working on improving the connection handling with the upstream - a welcoming news.

Distributed transactions are implemented using [Two-Phase commit protocol](https://martinfowler.com/articles/patterns-of-distributed-systems/two-phase-commit.html). Citus uses [pg_auto_failover](https://github.com/citusdata/pg_auto_failover) extension for implementing HA.

The last part of the paper is about benchmarks. It seems like Citus is winning in most of the scenarios. I generally take benchmarks with a pinch of salt.

#### Open Questions:
1. What's the difference (in features) between open source Citus and paid?
   - Nothing. [While I was reading the paper](https://twitter.com/viggy28/status/1537665649241075712), Citus completely [open sourced all their enterprise features](https://github.com/citusdata/citus/pull/6008). Impressed with their work. ![Citus open source](/images/citus-tweet-2.png)
2. Does each tenant have its own shard (aka table) in a worker node?
3. How does it manages things like DDL change, sequences, truncate?

#### References:
 - [v11 release notes](https://www.citusdata.com/updates/v11-0/)
 - https://www.youtube.com/watch?v=JwjjUT8K7po