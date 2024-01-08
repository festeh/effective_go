package main

import (
	"fmt"
	"os"
)

func main() {
	flags := &Flags{
		nRequests: 10,
		nThreads:  5,
	}
	err := ParseCLI(flags)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Running %d requests to %q with %d threads\n", flags.nRequests, flags.URL, flags.nThreads)
}
