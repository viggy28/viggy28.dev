---
date: 2019-04-09T20:00:08-04:00
description: "Deploying cloud run using Gitlab"
featured_image: ""
tags: ["cloud run", "gitlab", "docker"]
title: "How to deploy a cloud run application using Gitlab CI/CD"
---

In Google Cloud Next 2019, they introduced a new product called [Cloud Run] (<https://cloud.google.com/run/>). I've been using it from the EAP (Early Access Program) days. [As the name suggests,  basically it runs your docker image. You might be familiar with other serverless products such as [Cloud Function] (<https://cloud.google.com/functions/>) or [lambda] (<https://aws.amazon.com/lambda/>) where you provide your source code instead of a docker image. In my opinion, Cloud Run is more flexible than functions. Let me explain why.

+ You can write your application in any language and any version you wish
  
+ You can easily migrate it from one provider to another. Docker is **THE** standard for running container workloads.

+ Don't have to deal with handler or file names or zipping source code.

So far so good. What happens if I need to make a code change. Easy right? create a Docker image, push it to Google Container Registry and then create a new cloud run service. Imagine doing this manually every time.

This article is about how to automate those steps using [Gitlab CI/CD] (<https://docs.gitlab.com/ce/ci/>). The gitlab repository is available [here] (<https://gitlab.com/viggy28-websites/code-samples/cloudrun-cicd>)

Let's deploy that [go application] (<https://gitlab.com/viggy28-websites/code-samples/cloudrun-cicd/blob/master/main.go>) to cloud run.

[Dockerfile] (<https://gitlab.com/viggy28-websites/code-samples/cloudrun-cicd/blob/master/Dockerfile>) to build a Docker image.

```Dockerfile

# Use the official Golang alpine image to create a build artifact.
# https://hub.docker.com/_/golang
FROM golang:alpine as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/viggy28/cloudrun-cicd
COPY . .

# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN apk add git \
&& go get github.com/sirupsen/logrus

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o cloudrun-cicd

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/viggy28/cloudrun-cicd/cloudrun-cicd /cloudrun-cicd

# Run the web service on container startup.
CMD ["/cloudrun-cicd"]

```

Note:

+ Using golang:alpine docker image as the base
+ Doing go get of the external package
+ To minimize the size of the docker image, We are doing multi-stage builds

Up next is the [.gitlab-ci.yml] (<https://gitlab.com/viggy28-websites/code-samples/cloudrun-cicd/blob/master/.gitlab-ci.yml>) file.

```yml
stages:
  - build-push-docker
  - deploy-new-service

variables:
    projectid: "cloudrundemo-237205"
    service: "cloudrun-cicd"
    region: "us-central1"
    image_name: "cloudrun-cicd"

build-push-docker-job:
  stage: build-push-docker
  image: docker:latest
  services:
    - docker:dind
  before_script:
    - echo $CLOUDRUN_CICD_SA_KEY > ${HOME}/service_key.json
    - cat ${HOME}/service_key.json | docker login -u _json_key --password-stdin https://gcr.io
  script:
    - docker build -t gcr.io/$projectid/$image_name:latest .
    - docker push gcr.io/$projectid/$image_name:latest

deploy-new-service-job:
  stage: deploy-new-service
  before_script:
    # Download and install Google Cloud SDK. Source it so you can use it from any where
    - mkdir /software
    - cd /software
    - wget https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz
    - tar zxvf google-cloud-sdk.tar.gz && ./google-cloud-sdk/install.sh --usage-reporting=false --path-update=true
    - source /root/.bashrc
    - gcloud components install beta
    # Write our GCP service account private key into a file
    - echo $CLOUDRUN_CICD_SA_KEY > ${HOME}/service_key.json
  script:
    - gcloud auth activate-service-account --key-file ${HOME}/service_key.json
    - gcloud beta run deploy $service --region $region --project $projectid --image gcr.io/$projectid/$image_name:latest
    - gcloud beta run services add-iam-policy-binding $service --project $projectid --member="allUsers" --role="roles/run.invoker" --region $region
```

**Prerequisites**:

+ Create a Service Account (SA) in GCP. I prefer to have a dedicated SA for CI/CD. (Of course, you can use your existing one)
  
+ Grant "Service Account User", "Cloud Run Admin" and "Storage Admin" roles to the newly created.
  
    ![SA](/images/cloudrun-cicd-iam.png)

+ Download the SA JSON key and copy it
  
+ Create an environment variable in Gitlab CI/CD with the JSON key as the value. CLOUDRUN_CICD_SA_KEY is an environment variable which contains the key value copied in the above step

That should be it. Just make changes to your code and push it to gitlab. build-push-docker and deploy-new-service stages should trigger automatically every time when you push the code 

![Your CI/CD job should look something like this](/images/cloudrun-cicd-jobstatus.png)