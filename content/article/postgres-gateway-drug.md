---
date: 2026-03-14T08:00:08-04:00
description: "Every data platform now starts at Postgres — because that's where the data starts"
featured_image: "/images/pg-gateway.png"
tags: ["postgres", "data warehouse"]
title: "Postgres Is the Gateway Drug"
---

*The analytics databases optimized for queries. The real leverage was always where data gets written.*

---

Your application writes to Postgres. So does almost everyone else's. Postgres is the most widely used database in the world, [topping Stack Overflow's developer survey](https://survey.stackoverflow.co/2024/technology#most-popular-technologies-database) two years running. It didn't win as a data warehouse. It won as the place applications write.

For years, the data infrastructure companies competed on what happens after — faster queries, better compression, smarter optimizers. The traditional model relied on storage gravity: move your data into our walls, pay us to store and query it. Once data lived inside a warehouse, the switching cost kept it there.

Apache Iceberg removed that gravity.

Iceberg is an open table format — Parquet files plus metadata manifests, with a spec any engine can implement. When data lives in Iceberg on S3, it isn't in anyone's warehouse. Snowflake can query it. Databricks can query it. DuckDB can query it. The storage layer is commoditized.

So if data doesn't have to live in their storage — and a Postgres extension can write directly to open-format object storage without touching their compute — what exactly are they selling? The query engine. The governance layer. The catalog. The lock-in has moved up the stack.

The race is now about something different: who is closest to the origin of the data, and who writes it into the format everything else reads. That origin is Postgres.

![before and now](/images/pg-gateway-drug.png)

---

## Follow the Money

Don't take my word for it — look where over $1.5 billion went in eighteen months.

[ClickHouse acquired PeerDB](https://clickhouse.com/blog/clickhouse-acquires-peerdb-to-boost-real-time-analytics-with-postgres-cdc-integration) (July 2024) — the fastest Postgres CDC (change data capture) tool — and built it into their ingestion layer. Then they [launched a managed Postgres service](https://clickhouse.com/blog/postgres-managed-by-clickhouse) (January 2026). The pipeline from Postgres wasn't enough. They wanted to own the instance.

[Databricks acquired Neon](https://www.databricks.com/company/newsroom/press-releases/databricks-agrees-acquire-neon-help-developers-deliver-ai-systems) for ~$1B (May 2025) — a serverless Postgres platform — and launched [Lakebase](https://www.databricks.com/blog/announcing-lakebase-public-preview): Postgres integrated directly into their Lakehouse, with no ETL between operational data and analytics. Four months later they [acquired Mooncake Labs](https://www.databricks.com/blog/mooncake-labs-joins-databricks-accelerate-vision-lakebase) (October 2025), whose open source tools mirror Postgres WAL (write-ahead log) changes directly into Iceberg and Delta Lake in real time.

[Snowflake acquired Crunchy Data](https://www.snowflake.com/en/blog/snowflake-postgres-enterprise-ai-database/) for ~$250M (June 2025). [Crunchy's own announcement](https://www.crunchydata.com/blog/crunchy-data-joins-snowflake) described why: customers wanted to work at "the intersection of Postgres and Apache Iceberg." In October 2025, Snowflake shipped [write support for externally managed Iceberg tables](https://www.snowflake.com/en/blog/external-iceberg-tables-generally-available/) — data living entirely outside their storage. Then in February 2026, they open sourced [`pg_lake`](https://www.snowflake.com/en/engineering-blog/pg-lake-postgres-lakehouse-integration/), letting Postgres write to Iceberg without Snowflake's compute layer involved at all.

Three companies. Same bet. Own the Postgres layer.

---

## The Open Source Signal

The acquisitions are a capital signal. The open source projects are an engineering signal — and in some ways the cleaner one.

All three companies independently built and published a Postgres-to-Iceberg bridge:

| Project | Author | What it does |
|---|---|---|
| **[PeerDB](https://github.com/PeerDB-io/peerdb)** | ClickHouse | Postgres CDC engine, open source, powers ClickPipes |
| **[`pg_mooncake` + Moonlink](https://github.com/Mooncake-Labs/pg_mooncake)** | Databricks (via Mooncake Labs) | Postgres WAL → Iceberg/Delta Lake, columnar queries inside Postgres |
| **[`pg_lake`](https://github.com/Snowflake-Labs/pg_lake)** | Snowflake (via Crunchy Data) | Postgres → Apache Iceberg, native open format read/write bridge |

When three competitors converge on the same open source architecture independently, that's not a coincidence. It's a structural shift.

---

## What This Means

The data warehouse isn't going away — but its role is shrinking. When storage lives in open formats on object storage and Postgres can write to it directly, the warehouse becomes one query engine among many. The moat moves from "we hold your data" to "we govern and optimize your queries."

For data engineers, the center of gravity shifts left. The most consequential infrastructure decisions won't be which warehouse to use — they'll be how data leaves Postgres. Extensions like `pg_mooncake` and `pg_lake` turn Postgres into the first mile of the data platform, not just the application backend.

The companies that understood this earliest spent $1.5 billion to own that first mile. The rest of the industry will follow.

---

*Timeline: [ClickHouse acquires PeerDB](https://clickhouse.com/blog/clickhouse-acquires-peerdb-to-boost-real-time-analytics-with-postgres-cdc-integration) (July 2024) · [Databricks acquires Neon](https://www.databricks.com/company/newsroom/press-releases/databricks-agrees-acquire-neon-help-developers-deliver-ai-systems) (May 2025) · [Snowflake acquires Crunchy Data](https://www.snowflake.com/en/blog/snowflake-postgres-enterprise-ai-database/) (June 2025) · [Databricks acquires Mooncake Labs](https://www.databricks.com/blog/mooncake-labs-joins-databricks-accelerate-vision-lakebase) (October 2025) · [ClickHouse launches native Postgres](https://clickhouse.com/blog/postgres-managed-by-clickhouse) (January 2026) · [Snowflake Postgres GA + `pg_lake` open sourced](https://www.snowflake.com/en/engineering-blog/pg-lake-postgres-lakehouse-integration/) (February 2026)*