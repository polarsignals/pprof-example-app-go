VERSION:=$(shell cat VERSION | tr -d '\n')
CONTAINER_IMAGE:=quay.io/polarsignals/pprof-example-app-go:$(VERSION)

LDFLAGS="-X main.version=$(VERSION)"

all: pprof-example-app-go pprof-example-app-go-stripped pprof-example-app-go-withput-dwarf pprof-example-app-go-inlining-disabled
pprof-example-app-go: go.mod main.go fib/fib.go
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -o $@ main.go

.PHONY: container
container: pprof-example-app-go
	docker build -t $(CONTAINER_IMAGE) .

pprof-example-app-go-stripped: go.mod main.go fib/fib.go
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -ldflags="-s -w" -trimpath -o $@ main.go

pprof-example-app-go-without-dwarf: go.mod main.go fib/fib.go
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -ldflags="-w" -o $@ main.go

pprof-example-app-go-inlining-disabled: go.mod main.go fib/fib.go
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -gcflags="all=-N -l" -o $@ main.go
