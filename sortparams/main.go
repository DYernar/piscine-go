package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for j := 0; j < len(args)-1; j++ {
		for i := 0; i < len(args)-1; i++ {
			if args[i] > args[i+1] {
				temp := args[i]
				args[i] = args[i+1]
				args[i+1] = temp
			}
		}
	}
	for index, item := range args {
		if index != 0 {
			for _, itemOfItem := range item {
				z01.PrintRune(itemOfItem)
			}
			z01.PrintRune('\n')
		}
	}

}
