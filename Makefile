VERSION:=$(shell cat VERSION | tr -d '\n')
CONTAINER_IMAGE:=quay.io/polarsignals/pprof-example-app-go:$(VERSION)

LDFLAGS="-X main.version=$(VERSION)"

pprof-example-app-go: go.mod main.go fib/fib.go
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -o $@ --installsuffix cgo main.go

.PHONY: container
container: pprof-example-app-go
	docker build -t $(CONTAINER_IMAGE) .
