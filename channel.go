package main

import (
	"fmt"
	"time"
)

func out(from,to int,ch chan bool)  {
	for i:=from;i<=to;i++{
		time.Sleep(5000 *time.Millisecond)
		fmt.Println(i)
	} 
	ch<- true
}
func main()  {
	ch:=make(chan bool)
	go out(0,5,ch)
	go out(6,10,ch)
	<-ch

	
	
}