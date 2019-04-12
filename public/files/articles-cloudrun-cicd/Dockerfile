# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
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