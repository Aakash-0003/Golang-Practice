package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("Hello world \n")
	sli := "hello"
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%", sli[i])
	}
	//fmt.Println("Default value of GOMAXPROCS is: ", runtime.NumCPU())
}
