package main

import (
	"fmt"
	"os"

	"github.com/Neal-C/gh-ghpm/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
