package main

import (
	"os"

	"github.com/apzelos/zw/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
