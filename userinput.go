package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	text :="Write name"
	fmt.Println(text)
	reader :=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fmt.Println("Welcome ",input)	

	//type conversion
	text1 :="Write number"
	fmt.Println(text1)
	reader1 :=bufio.NewReader(os.Stdin)
	input1,_:=reader1.ReadString('\n')
	number,err:=strconv.ParseFloat(strings.TrimSpace(input1),64)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Number +1 =",number+1 )
	}
	

	
}