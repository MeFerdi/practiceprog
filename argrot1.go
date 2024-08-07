package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}

	Args := os.Args[1:]

	rotatedArgs := make([]string, len(Args))
	for i := 0; i < len(Args); i++ {
		rotatedArgs[i] = Args[(i+1)%len(Args)]
	}

	for _, arg := range rotatedArgs {
		fmt.Print(arg + " ")
	}
	fmt.Println()
}
