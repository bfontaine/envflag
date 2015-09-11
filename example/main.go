package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bfontaine/envflag"
)

func main() {
	envflag.AutoSetup()
	flag.Parse()

	fmt.Printf("The key is '%s'\n", os.Getenv("KEY"))
}
