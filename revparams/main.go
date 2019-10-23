package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	length := 0
	for index := range os.Args {
		length = index
	}
	for i := length; i > 0; i-- {
		for _, item := range os.Args[i] {
			z01.PrintRune(item)
		}
		z01.PrintRune('\n')
	}
}
