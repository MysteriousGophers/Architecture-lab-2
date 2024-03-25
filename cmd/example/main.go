package main

import (
	"errors"
	"flag"
	"fmt"
	lab2 "github.com/mysteriousgophers/architecture-lab-2"
	"io"
	"os"
	"strings"
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

	reader, writer, closer, err := getReaderAndWriter()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	defer func() {
		if closer != nil {
			if err := closer(); err != nil {
				fmt.Fprintln(os.Stderr, "Error closing file: ", err)
			}
		}
	}()
	handler := &lab2.ComputeHandler{Reader: reader, Writer: writer}
	err = handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func getReaderAndWriter() (io.Reader, io.Writer, func() error, error) {
	if inputExpression == "" && inputFile == "" {
		return nil, nil, nil, errors.New("no provided expression")
	} else if inputExpression != "" && inputFile != "" {
		return nil, nil, nil, errors.New("provided expression in both ways")
	}
	var reader io.Reader
	var writer io.Writer
	var closer func() error
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, nil, nil, errors.New("error reading file")
		}
		reader = file
	} else {
		reader = strings.NewReader(inputExpression)
	}
	if outputFile != "" {
		file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, nil, nil, errors.New("error opening file to write output to")
		}
		writer = file
		closer = func() error {
			return file.Close()
		}
	} else {
		writer = os.Stdout
		closer = nil
	}
	return reader, writer, closer, nil
}
