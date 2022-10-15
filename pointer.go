package main

import "fmt"
func main()  {
	number := 10
	pointer := &number
	fmt.Println("memort adress of Number:",pointer)
	fmt.Println("Number:",*pointer)
}