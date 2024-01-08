package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	URL       string
	nRequests int
	nThreads  int
}

func Validate(flags *Flags) error {
	if flags.nRequests < 1 {
		return fmt.Errorf("Invalid number of requests: %d", flags.nRequests)
	}
	if flags.nThreads < 1 {
		return fmt.Errorf("Invalid number of threads: %d", flags.nThreads)
	}
	if flags.URL == "" {
		return fmt.Errorf("Invalid URL: %q", flags.URL)
	}
	return nil
}

func ParseCLI(flags *Flags) error {
	flag.IntVar(&flags.nRequests, "n", flags.nRequests, "Number of requests to perform")
	flag.IntVar(&flags.nThreads, "c", flags.nThreads, "Number of multiple requests to make at a time")
	flag.StringVar(&flags.URL, "u", flags.URL, "URL to hit")

	flag.Parse()
	return Validate(flags)
}
