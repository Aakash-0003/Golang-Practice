package main

import (
	"fmt"
	"sync"

	"time"
)

var wg sync.WaitGroup

func timer(count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println("ughh")

	}
}
func timer2(count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println("yeahh")

	}
}

func main() {
	start := time.Now()
	go timer(3)
	wg.Add(1)
	go timer2(3)
	wg.Add(1)
	wg.Wait()
	fmt.Println(time.Since(start))

}
