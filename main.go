package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var THREADS int
var URL string
var TIMEOUT int

var TOTALHITREQUEST int

var respCount map[int]int
var timeTaken []float64

var mux sync.Mutex

var choices = []byte{'-', '\\', '|', '/'}

func init() {
	flag.StringVar(&URL, "u", "https://google.com", "URL")
	flag.IntVar(&THREADS, "t", 4, "Number of threads.")
	flag.IntVar(&TIMEOUT, "s", 10, "Time to Run")
	flag.Parse()
	fmt.Printf("%10s: %10s\n", "URL", URL)
	fmt.Printf("%10s: %10d\n", "Threads", THREADS)
	fmt.Printf("%10s: %10d\n", "Time", TIMEOUT)
}

func requester() {

	for {

		start := time.Now()
		resp, _ := http.Get(URL)
		end := time.Since(start)
		mux.Lock()
		if resp != nil {
			respCount[resp.StatusCode]++
		} else {
			respCount[0]++
		}
		mux.Unlock()
		TOTALHITREQUEST++

		timeTaken = append(timeTaken, end.Seconds())

		fmt.Printf("[%c]: Processing...\r", choices[TOTALHITREQUEST%4])
	}

}

func main() {

	timeTaken = make([]float64, 0)
	respCount = make(map[int]int)
	wg.Add(1)
	for i := 0; i < THREADS; i++ {
		go requester()
	}

	timer := time.NewTimer(time.Duration(TIMEOUT) * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Print("\r\n")
			fmt.Println(strings.Repeat("-", 30))
			fmt.Printf("|%15s | %10s|\n", "Status Code", "Hits")
			fmt.Println(strings.Repeat("-", 30))
			for k, v := range respCount {
				fmt.Printf("|%15d | %10d|\n", k, v)
			}
			fmt.Println(strings.Repeat("-", 30))
			var maxTime, minTime, avgTime float64
			minTime = timeTaken[0]
			for _, v := range timeTaken {
				if v > maxTime {
					maxTime = v
				} else if v < minTime {
					minTime = v
				}
				avgTime += v
			}
			avgTime = avgTime / float64(len(timeTaken))

			fmt.Println(strings.Repeat("-", 30))

			fmt.Printf("|%-15s | %10d|\n", "Total Hits", TOTALHITREQUEST)
			fmt.Println(strings.Repeat("-", 30))
			fmt.Printf("|%-15s | %10.2f|\n", "Max Time", maxTime)
			fmt.Println(strings.Repeat("-", 30))
			fmt.Printf("|%-15s | %10.2f|\n", "Min Time", minTime)
			fmt.Println(strings.Repeat("-", 30))
			fmt.Printf("|%-15s | %10.2f|\n", "Average Time", avgTime)
			fmt.Println(strings.Repeat("-", 30))
			return

		}
	}
}
