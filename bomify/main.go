package main

import (
	"bytes"
	"io"
	"os"
)

var bom = []byte{0xef, 0xbb, 0xbf}

func main() {
	io.Copy(os.Stdout, io.MultiReader(bytes.NewReader(bom), os.Stdin))
}
