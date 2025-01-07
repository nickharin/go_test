package go_test

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	fmt.Println("Hello World \n Привет катюха")
}
