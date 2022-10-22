package main

import (
	"fmt"
)

func main() {
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var i string
	fmt.Scanln(&i)
	results = append(results, i)
	res := 0
    for _, v := range results {
        switch v {
            case "w":
            res += 3
            case "l":
            res += 0
            case "d":
            res += 1
        }
    }
    fmt.Println(res)
  
}
    
    
