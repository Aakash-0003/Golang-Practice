package main

import (
	"fmt"
	"runtime"
)

func main() {
	var input int
	fmt.Scan(&input)
	output := int32(input)
	fmt.Println(output)
	fmt.Println(runtime.NumCPU())
}
