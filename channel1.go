package main

import (
	"fmt"
	"time"
)
func evenSum(from, to int, ch chan int) {

	result := 0
  
	for i:=from; i<=to; i++ {
		time.Sleep(500 *time.Millisecond)
  
	  if i%2 == 0 {
  
		result += i
  
	  }    
  
	}
  
	ch <- result
  
  }
  
  func squareSum(from, to int, ch chan int) {
  
	result := 0
  
	for i:=from; i<=to; i++ {
  
	  if i%2 == 0 {
  
		result += i*i
  
	  }    
  
	}
  
	ch <- result
  
  }
  func main()  {
	sqch := make(chan int)
	evch := make(chan int)
	go evenSum(1, 10, evch)
	go squareSum(1, 10, sqch)
	fmt.Println(<-evch + <-sqch)
	
  }