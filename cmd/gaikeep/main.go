package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/christimiller/gaikeep/internal/gaikeep"
)

func main() {
	flag.Parse() // Parse command-line arguments

	if err := gaikeep.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
