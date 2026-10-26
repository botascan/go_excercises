package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("fastax", flag.ContinueOnError)
	input := flags.String("input", "", "input file")
	output := flags.String("output", "", "output file")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if *input == "" {
		return errors.New("-input is required")
	}

	fastaseq, err := os.ReadFile(*input)
	if err != nil {
		return err
	}

}
