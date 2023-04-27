---
date: 2023-04-26T10:40:00-07:00
description: "Lessons learnt from building a managed Postgres offering using Open Source tools"
featured_image: ""
tags: ["postgres"]
title: "Challenges of Building in-house RDS"
---

**Abstract**

With the advent of cloud/managed offerings, running Postgres on their own hardware is becoming rare. It's great that a lot of the Ops work is taken care of by the provider, however, understanding a layer or two beneath these abstractions will be useful for anyone in strengthening their knowledge.

For eg. When you can't connect to an instance how do you find where the problem is?

a. Are you able to reach the server i.e. Is ICMP even working? 

b. Is the service listening on the right interface and IP? 

c. Are the IPtables set up correctly to allow communication? 

I found these problems more prevail in on-prem setup.


What happens if all of sudden the DMLs take longer to finish? There is no performance insight

How to respond if the server runs out of disk? There is no EBS. 

How to manage if there is a sudden spike in database traffic? There is no Auto Scaling Group

 

Also, realize that one can't just raise a customer support ticket and wait for a solution. Engineers need to role-up their sleeves and find the root cause

There are bigger challenges.

What happens if a server goes down?

What happens if a data center goes down?

What happens if a region goes down?


These are not hypothetical questions if you are running on-prem. These things happen more frequently than one imagines. 


This talk tries to answer some of the above questions based on the experience that I gained at Cloudflare. Also, hopefully, piques the interest of the audience to give it a try on on-prem if they are looking to do it.

[Slides for the presentation](https://docs.google.com/presentation/d/1SbkgBdHK0vcGTIQT_C2b-WxUoQyGU4-9fH3OzhRY_no/edit?usp=sharing)

[Recordings of the presentation](https://www.youtube.com/watch?v=cBVRtrEhbDo)

[Link from SCaLE website](https://www.socallinuxexpo.org/scale/20x/presentations/challenges-building-house-rds)


