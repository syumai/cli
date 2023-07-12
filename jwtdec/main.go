package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if err := runMain(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var ErrInvalidArgument = errors.New("Invalid Argument")

func runMain() error {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	sp := strings.Split(string(b), ".")
	if len(sp) != 3 {
		fmt.Printf("split length must be 3, got: %d", len(sp))
		return ErrInvalidArgument
	}
	sp = sp[0:2]
	for _, s := range sp {
		b, err := base64.RawURLEncoding.DecodeString(s)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}
	return nil
}
