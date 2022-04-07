package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(str)
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Printf("FOUND")
	} else {
		fmt.Printf("NOT FOUND")
	}

}
