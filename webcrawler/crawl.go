package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: crawl <link>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		usage()
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	queue := make(chan String)
	filteredQueue := make(chan String)

	go func() { queue <- args[0] }()
	go filterQueue(queue, filteredQueue)

	done := make(chan bool)

	for i := 0; i < 5; i++ {
		go func() {
			for uri := range filteredQueue {
				enqueue(uri, queue)
			}
			done <- true
		}()
	}
	<-done
}

func filterQueue(in chan string, out chan string) {
	var seen = make(map[string]bool)
	for val := range in {
		if !seen[val] {
			seen[val] = true
			out <- val
		}
	}
}
