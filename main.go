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

const (
	modeBusyCPU = "busyCPU"
	modeAllocMem = "allocMem"
)

func main() {
	var (
		bind = ""
		mode = "both" // busyCpu, allocMem
	)
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.StringVar(&bind, "bind", ":8080", "The socket to bind to.")
	flagset.StringVar(&mode, "mode", "both", "The mode to run. Options: busyCPU, allocMem, both")
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

	switch mode {
	case modeBusyCPU:
		go busyCPU()
	case modeAllocMem:
		go allocMem()
	default:
		go busyCPU()
		go allocMem()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(c)
	<-c
}

// Calculates Fibonacci numbers starting with 1 000 000th.
// Produces some CPU activity.
func busyCPU() {
	i := uint(1000000)
	for {
		log.Println("fibonacci number", i, fib.Fibonacci(i))
		i++
	}
}

// Allocate 1mb of memory every second, and don't free it.
// Don't do this at home.
func allocMem() {
	buf := []byte{}
	mb := 1024 * 1024

	for {
		buf = append(buf, make([]byte, mb)...)
		log.Println("total allocated memory", len(buf))
		time.Sleep(time.Second)
	}
}
