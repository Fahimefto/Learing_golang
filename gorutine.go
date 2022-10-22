package main

import (
	"fmt"
	"time"
)

func out(from,to int)  {
	for i:=from;i<to;i++{
		time.Sleep(500 *time.Millisecond)
		fmt.Println(i)
	} 
}
func main()  {
	go out(1,10)
	go out(2,5)

	time.Sleep(5 *time.Second)
	
}