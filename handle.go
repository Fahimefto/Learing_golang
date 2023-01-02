package main

import (
	"fmt"
	"net/url"
)

var web string="https://technovent.sust.edu/"
func main() {
	result, err := url.Parse(web)
	if err != nil {
		panic(err)}

	fmt.Println(result)
	fmt.Println(result.Hostname())
}