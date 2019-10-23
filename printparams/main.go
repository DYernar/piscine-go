package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for index, item := range os.Args {
		if index != 0 {
			for _, itemsInItem := range item {
				z01.PrintRune(itemsInItem)
			}
			z01.PrintRune('\n')
		}
	}
}
