package main

import "fmt"
func main()  {
	var fruitlist1= [7]string{"apple","orange","banana","grapes","mango","pineapple","watermelon"}
	 var s []string = fruitlist1[2:5]
	fmt.Println(s)

	//// make 
	a:=make([]int,5)

	a= append(a, 1,2,3,4,5)
	fmt.Println(a)
	for _,v := range a{
		fmt.Println(v)
		
	}
}