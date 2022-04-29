package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

const (
	// Binary, See: http://en.wikipedia.org/wiki/Binary_prefix
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB
	PiB = 1024 * TiB
)

type Person interface {
	Name() string
	Live()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	// Limit the CPU usage to avoid overloading.
	runtime.GOMAXPROCS(1)
	// Set sampling rate for mutex and block profile
	// See: https://pkg.go.dev/net/http/pprof
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	pp := []Person{&alice{}, &bob{}, &dave{}, &eve{}}

	for {
		for _, p := range pp {
			p.Live()
		}
		time.Sleep(time.Second)
	}
}
