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

Install the package requests using pip to execute this program locally.

Next, I wrote a Dockerfile which basically bundles base image, commands to execute and install other packages.

Then build a docker image by executing docker build -t <imageName> . You can verify the image by running docker images. (note down the image id) To test the image is working as expected docker run <imageName>

Didn’t want to stop there. How cool it will be to run the image in a different server!!?.

Alright, you need to push the docker image which you build locally to the docker hub, so that you can pull and run the image from different servers.

Preparation to push the docker image.

```
docker tag <imageid><dockerHubUserName>/<imageName:latest>

docker login ( provide dockerHubUserName/Password)

docker push <dockerHubUserName>/<imageName>
```

Successfully, you have pushed a docker image to the docker hub !!.

To really appreciate the use of containerization, you need to run this image from a different environment. Hmm!!? The easiest way is if you have a VM, you can use that (or a different computer). If you don’t have both, then you can spin up a VM from Google Cloud. (or from any other provider).

Note : There are OS images which includes Docker. So, you don’t even have to install Docker in the VM.

Boot Disk : Container-Optimized OS 61–9765.66.0 stable

Kernel: ChromiumOS-4.4.70 Kubernetes: 1.6.10 Docker: 17.03.2

{{< figure src="/images/gcp.png" title="Spinning up a Google Compute Engine (VM)" >}}

Once you created the VM, ssh to the box. You should see on the dashboard. Then follow the commands in the below screenshot.

{{< figure src="/images/bash.png" title="Running a docker image from Google Compute Engine" >}}

You don’t have to install Python, no need to install the package (requests). Everything is bundled in the image and just execute it !!

[Source Code] (<https://github.com/vr001/python_docker>)

[Docker Image] (<https://hub.docker.com/r/vira28/python_docker/>)

Reference :

[Requests: HTTP for Humans - Requests 2.18.4 documentation] (<http://www.python-requests.org/en/master/>)

[Dockerize your Python Application] (<https://runnable.com/docker/python/dockerize-your-python-application>)

