package main

import (
	"flag"
	"fmt"
	lab2 "go_lab_second"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Path/to/expression-file")
	outputFile      = flag.String("o", "", "Path/to/output-file")
)

func main() {
	flag.Parse()

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot open file.", err)
		}

		defer file.Close()

		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Usage: go run cmd/example/main.go -e <expression> | -f <path/to/expression-file [-o <path/to/outputfile>]")
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot create file.", err)
		}

		defer file.Close()

		output = file
	} else {
		output = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:      input,
		Output:     output,
		Calculator: &lab2.PrefixCalculator{},
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Unexpected error occured.", err)
	}
}
