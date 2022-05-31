---
date: 2022-05-31
description: "Logging related postgres settings"
featured_image: ""
tags: ["postgres"]
title: "Postgres logging"
---

Before I forget let me write this down here.
$$
`log_stament`
    - none
    - all

`log_min_duration_statment`
    - millisecond value

When you set `log_statement`=none and `log_min_duration_statement`=1 then any statement which takes longer than 1 millisecond will be logged.

When you set `log_statement`=all and `log_min_duration_statement`=1 then all statements are logged; however it only shows duration on statements longer than 1 millisecond.

