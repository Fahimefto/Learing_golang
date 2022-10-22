package main

import (
	"fmt"
)
func evenSum(from, to int, ch chan int) {

	result := 0
  
	for i:=from; i<=to; i++ {
		
  
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
	go evenSum(1, 5, evch)
	go squareSum(6, 10, sqch)
	select{
	case <-evch:
		fmt.Println("evenSum is done")
		return
	
	}
	
	
  }