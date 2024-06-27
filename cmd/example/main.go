package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	data := new(strings.Builder)
	_, err := io.Copy(data, h.Input)
	if err != nil {
		return err
	}
	expression := strings.TrimSpace(data.String())
	result, err := PostfixToPrefix(expression)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(h.Output, result)
	return err
}

func main() {
	var expr string
	var inputFile string
	var outputFile string

	flag.StringVar(&expr, "e", "", "Вхідний постфіксний вираз")
	flag.StringVar(&inputFile, "f", "", "Файл з вхідним постфіксним виразом")
	flag.StringVar(&outputFile, "o", "", "Файл для запису префіксного виразу")
	flag.Parse()

	var input io.Reader
	var output io.Writer = os.Stdout

	if expr != "" {
		input = strings.NewReader(expr)
	} else if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Помилка відкриття файла:", err)
			return
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Необхідно вказати вхідний вираз або файл з виразом")
		return
	}

	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Помилка створення файла:", err)
			return
		}
		defer file.Close()
		output = file
	}

	handler := &ComputeHandler{Input: input, Output: output}
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Помилка обчислення:", err)
	}
}
