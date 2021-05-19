package main

import (
	"fmt"

	"github.com/rs/xid"
)

func main() {
	fmt.Println(xid.New())
}
