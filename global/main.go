package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/polarsignals/pprof-example-app-go/fib"
)

func main() {
	// Calculates Fibonacci numbers starting with 1 000 000th.
	// Produces some CPU activity.
	go calculateFib()

	// Allocate 1mb of memory every second, and don't free it.
	// Don't do this at home.
	go allocMem()

	log.Println(http.ListenAndServe(":8080", nil))
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
