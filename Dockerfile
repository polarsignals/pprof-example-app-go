FROM gcr.io/distroless/base

ADD pprof-example-app-go /bin/pprof-example-app-go

ENTRYPOINT ["/bin/pprof-example-app-go"]
