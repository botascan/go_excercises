package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const _name = "dnax"

func main() {

	if err := run(os.Args, os.Stdout); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error: %v", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flag := flag.NewFlagSet(_name, flag.ExitOnError)
	input := flag.String("input", "", "Input DNA sequence")
	output := flag.String("output", "", "Output analyse results (default to Stdout)")
	color := flag.Bool("color", false, "Colorize output")
	//flag.BoolVar(color, "color", false, "Use colorized output")

	if err := flag.Parse(args[1:]); err != nil {
		return err
	}
	if *input == "" {
		return errors.New("-input flag is required ")
	}

	if *color {
		enableColor()
	}

	byteSeq, err := os.ReadFile(*input)
	//fmt.Println(string(byteSeq[:len(byteSeq)-1]))
	seq := bytes.TrimSpace(byteSeq)

	if err != nil {
		return fmt.Errorf("input sequence file failed to open: %w", err)
	}

	result, err3 := analyseDNASeq(seq)

	if err3 != nil {
		return err3
	}

	var writer io.Writer
	if *output != "" {
		outFile, err := os.Create(*output)
		if err != nil {
			return fmt.Errorf("output file failed to open: %w", err)
		}
		defer outFile.Close()
		writer = outFile
	} else {
		writer = stdout
	}

	err2 := display(result, writer)
	if err2 != nil {
		return err2
	}
	return nil
}
