# Pprof Example App Go

This example app serves as an example of how one can easily instrument HTTP handlers with [Pprof][pprof] profiles in [Go](https://golang.org/). It uses the [Go standard library][gostdlib].

It calculates fibonacci numbers starting at number 1 million to produce some CPU activity and allocates memory 1mb per second, it never frees this memory so be careful it may crash your system.

> Run this example at your own risk Polar Signals is not responsible for damage this example may cause.

A Docker image is available at: `quay.io/polarsignals/pprof-example-app-go:v0.1.0`

## Deploying in a Kubernetes cluster

First, deploy an instance of this example application, which listens and exposes metrics on port 8080 using the following [Deployment manifest](manifests/deployment.yaml).

Then deploy the collector of the Polar Signals Continuous Profiler to continuously collect profiles from the example app.

[pprof]:https://github.com/google/pprof
[gostdlib]:https://golang.org/pkg/net/http/pprof/
