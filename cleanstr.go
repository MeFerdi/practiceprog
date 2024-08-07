package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	Args := os.Args[1]
	var result string
	space := false
	for i := 0; i < len(Args); i++ {
		str := Args[i]
		if str != ' ' || str != '\t' {
			result += string(str)
			space = false
		} else if !space {
			result += " "
			space = true
		}
	}
	if len(result) > 0 && result[len(result)-1] != ' ' {
		result = result[:len(result)-1]
	}
	fmt.Println(result)
}
