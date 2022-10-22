package main

import "fmt"
func main()  {
	a:=make([]int,5)

	a= append(a, 1,2,3,4,5)
	fmt.Println(a)
	for i,v := range a{
		fmt.Println(i,v)
	}

	res := 0

nums := [3]int{2, 4, 6}

for i, v := range nums {

  if i%2==0 {

    res += v

  }

}

fmt.Println(res)
}
	
