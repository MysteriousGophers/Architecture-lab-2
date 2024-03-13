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

	reader, writer, err := getReaderAndWriter()
	if err != nil {
		_, err1 := fmt.Fprintln(os.Stderr, "Error: ", err)
		if err1 != nil {
			fmt.Println(err1)
		}
	}
	handler := &lab2.ComputeHandler{Reader: reader, Writer: writer}
	err = handler.Compute()
	if err != nil {
		_, err1 := fmt.Fprintln(os.Stderr, "Error: ", err)
		if err1 != nil {
			fmt.Println(err1)
		}
	}
}

func getReaderAndWriter() (io.Reader, io.Writer, error) {
	if inputExpression == "" && inputFile == "" {
		return nil, nil, errors.New("no provided expression")
	} else if inputExpression != "" && inputFile != "" {
		return nil, nil, errors.New("provided expression in both ways")
	}
	var reader io.Reader
	var writer io.Writer
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			return nil, nil, errors.New("error reading file")
		}
		reader = file
	} else {
		reader = strings.NewReader(inputExpression)
	}
	if outputFile != "" {
		file, err := os.OpenFile(outputFile, os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, nil, errors.New("error opening file to write output to")
		}
		writer = file
	} else {
		writer = os.Stdout
	}
	return reader, writer, nil
}
