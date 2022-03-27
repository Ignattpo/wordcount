package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	count := len(strings.Fields(args[0]))
	fmt.Println(count)
}
