package main

import (
	"fmt"

	_ "github.com/sdboyer/deptest"
	"github.com/sigma/deptest/sub"
)

func init() {
	fmt.Printf("%s\n", sub.Version)
}

func main() {}
