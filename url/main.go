package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {

	res, err := http.Get(url)
	handleErr(err)
	//fmt.Printf("response : %T", res)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	defer fmt.Println("response status: ", res.Status)
	handleErr(err)
	fmt.Printf("Response from URL: %v \n", string(data))

}
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
