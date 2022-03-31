package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var el string
	sli := make([]int, 0)
	var x string = "y"
	for i := 0; ; i++ {
		if x == "x" {
			break
		} else {
			fmt.Println("enter element")
			fmt.Scan(&el)
			if el == "x" {
				break
			} else {
				input, _ := strconv.Atoi(el)
				if i < len(sli) {
					sli[i] = input
				} else {
					sli = append(sli, int(input))
				}
				sort.Ints(sli)
			}
		}
		fmt.Println("Sorted slice : ", sli)
		fmt.Println("enter x for exit")

	}
}
