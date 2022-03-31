package main

import (
	"fmt"

	"time"
)

func timer(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println("ughh")

	}
}
func timer2(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println("yeahh")

	}
}

func main() {
	start := time.Now()
	timer(3)
	timer2(3)
	fmt.Println(time.Since(start))
}
