package go_test

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	fmt.Print("Hello World")
}
