---
date: 2022-06-02T21:00:08-04:00
description: "A living doc on lessons learnt around postgres schema migrations"
featured_image: "/images/postgres-logo.png"
tags: ["postgres", "migration"]
title: "Postgres schema migration gotchas"
---

Capturing thoughts from https://twitter.com/viggy28/status/1530800893842444289

1. When you are doing major DML changes, other than locks one more thing to keep in mind is replication lag. Especially if you use your replicas in hot standby mode.

2. When you need to delete most of the records in a massive table, its better to create a new table and just copy the records that you need to preserve.
When you need to delete all the records in a massive table, just truncate it instead of deleting them.

3. Don't run the same migration from multiple sessions or pods. If your tooling is robust it won't happen in the first place.

4. More nuanced one. If you are using PgBouncer in transaction mode, then the settings you need on session-level won't be available. For eg. create index. Setting statement_timeout and running create index. Run it directly on Postgres.

References:
1. https://postgres.ai/blog/20220525-common-db-schema-change-mistakes
