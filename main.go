package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuender/fmove/move"
)

func main() {
	simulate := false

	flag.BoolVar(&simulate, "s", simulate, "simulate")
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		if err := move.Move(".", simulate); err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
		}

		return
	}

	for _, arg := range flag.Args() {
		if err := move.Move(arg, simulate); err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "fmove\n\n")
	fmt.Fprintf(os.Stderr, "File Move.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
