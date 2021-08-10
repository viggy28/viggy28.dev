---
date: 2021-08-09T23:20:08-05:00
description: "How to find whether an IP is in a CIDR or not"
featured_image: "/images/featured_image_networking.jpeg"
tags: ["networking"]
title: "IP in CIDR or not"
---

There are tools like https://tehnoblog.org/ip-tools/ip-address-in-cidr-range/ which does the job. It doesn't support IPv6. One of my colleagues today taught me how to do it by hand quickly.

Say we have an IP `2b06:4600:1101:0:abcd:efa:dbd:ea60:f5a6` is in the CIDR `2b06:4600:1101::/64`

First step is to check what's the mask bits. Here it is 64. So you take the address, take the first 64 bits and see if they are the same as the CIDR.

In this case, the first 64 bits are `2b06:4600:1101:0` which is the CIDR.

Each nibbles is 16 bit. If the mask is multiples of 16 then its easier. i.e. 48, 64, 32 etc whereas it is complicated if the mask is non multiples of 16. Say 24, 12