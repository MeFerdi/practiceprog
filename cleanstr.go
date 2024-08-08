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
	var result string
	Args := os.Args[1]
	space := false

	for i := 0; i < len(Args); i++ {
		char := Args[i]
		if char != ' ' && char != '\t' {
			result += string(char)
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
