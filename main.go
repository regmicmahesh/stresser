package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var THREADS int
var URL string
var TOTALHIT int

func init() {
	flag.StringVar(&URL, "u", "https://google.com", "URL")
	flag.IntVar(&THREADS, "t", 4, "Number of threads.")
	flag.IntVar(&TOTALHIT, "c", 100, "Number of hits.")
	flag.Parse()
}

func requester(count int) {

	defer wg.Done()

	for i := 0; i < count; i++ {
		http.Get(URL)
		fmt.Print(".")
	}

}

func main() {

	requestPerThread := TOTALHIT / THREADS

	for i := 0; i < THREADS; i++ {
		wg.Add(1)
		go requester(requestPerThread)
	}

	wg.Wait()

}
