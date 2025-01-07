package main

import (
	"fmt"
	"runtime"
)

func init() {
	if runtime.GOARCH == "386" {
		runtime.GOMAXPROCS(1)
	}
}

func main() {
	fmt.Print("Hello World")
}
