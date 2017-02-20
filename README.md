This repo hosts a tutorial on how to get started withi Google Container Engine (GKE).

# Slides

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