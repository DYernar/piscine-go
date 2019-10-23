package main

import (
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	length := 0
	for index := range args {
		length = index
	}

	for i := 1; i < length+1; i++ {
		if args[1] == "--upper" {
			val := piscine.Atoi(args[i])
			if val > 0 && val < 27 {
				newval := rune(piscine.Atoi(args[i]))
				z01.PrintRune(newval + 64)
			} else if i != 1 {
				z01.PrintRune(' ')
			}
		} else {
			val := piscine.Atoi(args[i])
			if val > 0 && val < 27 {
				newval := rune(piscine.Atoi(args[i]))
				z01.PrintRune(newval + 96)
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
