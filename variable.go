package main

import "fmt"

func main(){
	var name string="fahim"
	fmt.Println(name)
	fmt.Printf("type of name : %T \n",name)

	//implicit type declaration
	var age=20
	fmt.Println(age)

	//another way to declare variable
	age1:=28
	fmt.Println(age1)
}