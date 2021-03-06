This repo hosts a tutorial on how to get started with Google Container Engine (GKE).
It goes through setting up all the way to blue/green deployments using Kubernetes Deployments instead of the old replication controllers.

# Slides - Quick Link

[Go Talks Slides Link](http://go-talks.appspot.com/github.com/serinth/gke-example-go/slides/main.slide#1)

# Slides - Using Present Tool

Slides are using the [Golang Present tool](https://godoc.org/golang.org/x/tools/present) which can be installed into your GOPATH with:

```bash
go install golang.org/x/tools/cmd/present
```

Then to execute the slides:

```bash
cd slides
./present.sh
```

`present.sh` is a bash shorthand for executing the `main.slide` file with the present tool.

# Go Code

There are two main Go programs here that can also be executed via the slides.

1. API web application in `main.go` which hosts a health and a version endpoint
2. Watcher app which does an HTTP GET at a designated endpoint every 2 seconds and outputs the values

The *API* application is meant to be built into a docker container and then pushed to Google Container Registry (GCR)

The *watcher* application is meant for demonstrating Kubernetes' blue/green deployment (see the slides).

# Notes
This repo uses K8's deployments. The preferred way to manage these now is through a package manager such as [Helm](https://helm.sh/).
