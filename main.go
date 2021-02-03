package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/polarsignals/pprof-example-app-go/fib"
)

var (
	version string
)

func main() {
	bind := ""
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":8080", "The socket to bind to.")
	err := flagset.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting HTTP server on", bind)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	go func() { log.Fatal(http.ListenAndServe(bind, mux)) }()

	// Calculates Fibonacci numbers starting with 1 000 000th.
	// Produces some CPU activity.
	go calculateFib()

	// Allocate 1mb of memory every second, and don't free it.
	// Don't do this at home.
	go allocMem()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)
	<-c
}

func calculateFib() {
	i := uint(1000000)
	for {
		log.Println("fibonacci number", i, fib.Fibonacci(i))
		i++
	}
}

func allocMem() {
	buf := []byte{}
	mb := 1024 * 1024

	for {
		buf = append(buf, make([]byte, mb)...)
		time.Sleep(time.Second)
	}
}
