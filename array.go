package main

import "fmt"
func main()  {
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Orange"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))
	var fruitlist1= [6]string{"apple","orange","banana","grapes"}
	fmt.Println(fruitlist1)
	fmt.Println(len(fruitlist1))
	
}