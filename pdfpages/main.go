package main

import (
	"fmt"
	"os"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: pdfpages {filename}")
		return
	}
	num, err := pdf.PageCountFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(num)
}
