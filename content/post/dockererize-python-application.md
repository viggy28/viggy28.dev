---
date: 2017-09-23T18:00:08-04:00
description: "Dockerize python applicatio"
featured_image: "/images/docker-logo.png"
tags: ["docker", "python", "containers"]
title: "How to dockerize a python application"
---

### In this article we will see how to convert a simple python application to a containerized (Docker) one

Go to the profile of vignesh ravichandran
vignesh ravichandran
Sep 23, 2017
TL;DR -> Containerize a python app. Push the image and execute from Google Cloud.

On a leisurely Saturday afternoon, I thought its nice to play with containers.
I was fascinated with the idea of containerizing an app and running it across different environments. Its easier (for me) to understand something when I do, instead of just reading.

```
Requirements:

python ≥2.7 (and pip)

Docker ≥1.13.1

An account in docker hub and google cloud. (Both have free tier $)
```

This is the simple python program which I wanted to containerize.

```
import requests
city = “Los Angeles”
print “I am from %s” %city
r = requests.get(‘http://www.google.com')
print r.status_code
```

