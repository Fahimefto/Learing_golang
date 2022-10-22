package main

import "fmt"
func main()  {
	type Contact struct {
		name string
		age int
	}

	x := Contact{"James", 42}
	p := x
	fmt.Println(p.age)
}