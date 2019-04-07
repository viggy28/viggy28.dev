---
date: 2019-03-23T18:00:08-04:00
description: "Hosting a simple website"
featured_image: "/images/firebase-logo-wide.png"
tags: ["firebase", "hugo", "hosting"]
title: "How to host a simple static website"
---

### We all at some point in time thought about having a personal website. It may be for writing about your travel trips, posting pictures of food that you cooked or ranting about the tech which you don't like. I will leave the reason to you

## Step1: Buying a domain

[This] (<https://viggy28.dev>) websites domain name is **viggy28.dev**. You want a human friendly name, so that your friends, family and rest of the internet can reach your website.

You might've heard [GoDaddy] (<https://godaddy.com>). One of the famous ones. I personally use [Namecheap] (<https://namecheap.com>). I found their customer support much better than GoDaddy's.

## Step2: Hosting

Hosting the website is similar to hosting your relative. Basically, a place where you can host your website. Hosting can be expensive. Your domain registrar might provide hosting too.

There is a bunch of hosting providers. Find the comparison table by [pcmag] (<https://www.pcmag.com/roundup/316108/the-best-web-hosting-services>).

I am using [Firebase] (<https://firebase.com>) for few reasons.

1. My familiarity with the platform.

2. Its based of Google's infrastructure.

3. Ease of use (automatic SSL, quick deployments, other services such as firestore, authentication).

4. Its free tier is sufficient enough to run my site.

## Step3: Building the website

You can build it from the scratch by writing all the HTML/CSS. I am using [hugo] (<https://gohugo.io>) (a static site template builder) and [Ananke] (<https://github.com/budparr/gohugo-theme-ananke>) theme.

You can find the [source] (<https://gitlab.com/viggy28-websites/viggy28.dev>) of this website. It concludes my config, assets and themes.