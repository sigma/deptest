package main

import (
	"fmt"

	_ "github.com/sdboyer/deptest"
	"github.com/sigma/deptest/sub"
)

func init() {
	fmt.Println(sub.Version)
}

func main() {}
