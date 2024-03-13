package main

import (
	"flag"
	"fmt"
	lab2 "github.com/mysteriousgophers/architecture-lab-2"
	"io"
	"os"
)

var (
	inputExpression string
	inputFile       string
	outputFile      string
)

func main() {
	flag.StringVar(&inputExpression, "e", "", "Expression to compute")
	flag.StringVar(&inputFile, "f", "", "File to read from")
	flag.StringVar(&outputFile, "o", "", "File to write to")
	flag.Parse()

	reader, writer, err := getReaderAndWriter()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	handler := &lab2.ComputeHandler{Reader: reader, Writer: writer}
	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func getReaderAndWriter() (io.Reader, io.Writer, error) {
	return nil, nil, nil
}
