---
date: 2026-04-26T00:00:00-04:00
description: "The last line of defense for those OOPS moments — so even DBAs can sleep peacefully."
featured_image: ""
tags: ["postgres", "extensions", "databases"]
title: "pg_savior: a seatbelt for Postgres"
---

*The last line of defense for those OOPS moments — so even DBAs can sleep peacefully.*

Have you ever accidentally run a `DELETE` without a `WHERE` clause? Or typed `DROP TABLE` thinking you were in staging, only to realize a half-second too late that you were in prod?

If yes — keep reading, this is for you. If no — keep reading anyway. Nobody who uses Postgres long enough stays immune.

## The crown jewels

A few years ago I ran the Postgres team at Cloudflare. My manager at the time had a line he liked: "your team manages the crown jewels. Pretty much everything else here can be rebuilt — as long as we have this control-plane data."

That framing changed how I thought about Postgres. It isn't just a database. It's the substrate the business runs on. So the question becomes simple, and uncomfortable: if this is the most valuable thing we own, where exactly do we put the guard?

## Where the existing defenses leak

Most teams already have several guards. They're all good. They all leak.

**CI for migrations.** Every schema change goes through review and runs against a test database. Catches a lot. But:

- One-off migrations applied by hand never see CI.
- Hotfixes during incidents skip the playbook because speed beats process.
- CI runs against an empty seed database. A `DELETE FROM events WHERE tenant_id = $1` with `$1 = NULL` matches zero rows in test and every row in prod. Selectivity bugs are invisible to a test DB.
- The migration was correct — for staging. Got pointed at prod. CI can't catch a fat-finger on a connection string.

**Linters and code review for the application.** They see the SQL you *wrote*, not the SQL you *sent*. An ORM filter built from `request.user` silently drops the predicate when the user is anonymous, and `Model.objects.filter(user=None).delete()` clears the table. The diff that passed review and the query that runs in prod aren't always the same query.

**DBA process.** Posting your command in a shared channel for a second pair of eyes before you run it — even during an incident, especially during an incident. Gold standard. Still imperfect, because the people running these commands have admin shells. They are trusted by definition. The same humans you trust most have the fewest guardrails. And in the middle of a 3 a.m. page, the muscle memory that types `DELETE FROM jobs WHERE status = 'failed'` can just as easily type `DELETE FROM jobs;` and hit return.

The pattern across all three: every one of these defenses lives **upstream** of the database. They guard the path to the query. They don't guard the query itself.

## A different layer

You should have all of the above. CI on every migration. Linters on application code. A culture where DBAs paste commands into Slack for a second pair of eyes before running them, even mid-incident. These are necessary.

pg_savior is one more layer, sitting at the only place that sees every statement no matter who sent it: the database itself. psql, the application's ORM, a migration tool, a cron job, a support engineer's one-off script, an AI coding agent with database credentials — they all eventually hand a parse tree to Postgres. pg_savior hooks that step and refuses the obviously dangerous shapes.

```
postgres=# DELETE FROM emp;
ERROR:  pg_savior: DELETE without WHERE clause is blocked
HINT:  Add a WHERE clause, or set pg_savior.bypass = on for this session.
```

It catches the obvious shapes — `DELETE`/`UPDATE` without a `WHERE`, `CREATE INDEX` without `CONCURRENTLY`, `DROP DATABASE` — and the less obvious ones, like an `ALTER COLUMN TYPE` that quietly rewrites a billion-row table, or a `DELETE WHERE id > 0` whose planner estimate makes the intent clear. When you really do mean it, `SET LOCAL pg_savior.bypass = on` for the transaction and the guard steps aside. The full list of guards lives in the README; the point isn't to forbid, it's to make the destructive path require one extra deliberate keystroke.

## Why an extension, not a proxy

pg_savior is a Postgres extension — a shared library loaded into the Postgres process itself, not a proxy in front of it, not a sidecar, not a linter on the way in.

That distinction is the whole reason this works as a *last* line of defense. A proxy can be bypassed. A DBA shells into the host, opens psql against the local socket, and the proxy never sees the query. A linter only inspects what you ship; it never sees what an ad-hoc session types. Anything that can be routed around isn't a last line of defense — it's a suggestion.

The extension model puts pg_savior inside the planner itself. There is exactly one way into Postgres execution, and pg_savior is sitting on it.

## How it works

Three hooks, each doing one job:

1. **`post_parse_analyze_hook`** — fires after parse-analyze, before planning. If the statement is `DELETE`/`UPDATE` and `query->jointree->quals` is `NULL`, raise `ERROR`. Plan-shape independent, parameterized statements handled correctly, no planner work wasted on a query about to be refused.
2. **`ExecutorStart_hook`** — fires after planning, before execution. Reads the planner's row estimate from the `ModifyTable` node and refuses if it exceeds the configured threshold.
3. **`ProcessUtility_hook`** — fires for DDL. Catches the index, schema-rewrite, and drop cases.

Getting to those three was not the first try. Earlier versions walked the parse tree token-by-token and broke as soon as a `DELETE` used an indexscan + hashjoin. A second attempt walked the plan and was closer, but plan shape is too fluid to depend on. Settling on `post_parse_analyze_hook` made the check both simpler and correct across PG versions, including PG17 which had broken the older approach.

## Try it

```bash
make && sudo make install
```

Then add `shared_preload_libraries = 'pg_savior'` to `postgresql.conf`, restart, and `CREATE EXTENSION pg_savior;` in each database. Source and full config docs are at [github.com/viggy28/pg_savior](https://github.com/viggy28/pg_savior); the extension is also on PGXN.

Defense in depth wins. CI catches what it can see. Linters catch what they can parse. Process catches what humans remember to follow. pg_savior catches what makes it through all of those — so the next time you're woken at 3 a.m. for a Postgres issue, at least you don't have to wonder if you accidentally ran the wrong command. Sleep peacefully. The database is watching.
