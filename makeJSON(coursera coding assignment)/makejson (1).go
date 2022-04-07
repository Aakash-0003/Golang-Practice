package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, add string
	var m = make(map[string]string)
	fmt.Scan(&name)
	fmt.Scan(&add)
	m["address"] = add
	m["name"] = name
	fmt.Println(m)
	result, _ := json.Marshal(m)
	fmt.Printf(string(result))
}
