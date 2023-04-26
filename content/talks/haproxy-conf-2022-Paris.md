---
date: 2023-04-25T10:40:00-07:00
description: "Load balancing and high availability"
featured_image: ""
tags: ["postgres"]
title: "HAProxy Conference 2022 at Paris"
---

**Abstract**

This talk explores how Cloudflare uses HAProxy for health checks, load balancing and reading traffic among nodes set up with Postgres streaming replication in hot standby mode. Cloudflare operates multiple Postgres Clusters across four data centers, and all of these clusters are made up of six nodes.

During a primary failure, Cloudflare’s high availability system promotes a replica to become a primary, and HAProxy makes sure there is no write traffic between two primaries to avoid a split-brain scenario.

The presentation also explores how HAProxy helps improve Cloudflare’s security by encrypting traffic between data centers using SSL, ensuring no traffic across data centers is clear text.

[Slides for the presentation](https://docs.google.com/presentation/d/1NFaxHlnoWEcvUXqGfXOXpjL2lJDEWuCgGoSot11wvBU/edit?usp=sharing)

[Recordings of the presentation](https://www.youtube.com/watch?v=HIOo4j-Tiq4)

[Link from HAProxy website](https://www.haproxy.com/user-spotlight-series/load-balancing-and-high-availability-on-postgres-using-haproxy/)

[Blog on Cloudflare about performance isolation](https://blog.cloudflare.com/performance-isolation-in-a-multi-tenant-database-environment/)

