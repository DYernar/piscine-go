package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	var i int
	var isnegative bool = false
	for _, letter := range s {
		if letter == '-' {
			if isnegative || i != 0 {
				return 0
			} else {
				isnegative = true
			}
		} else if letter == '1' {
			i = (i * 10) + 1
		} else if letter == '2' {
			i = (i * 10) + 2
		} else if letter == '3' {
			i = (i * 10) + 3
		} else if letter == '4' {
			i = (i * 10) + 4
		} else if letter == '5' {
			i = (i * 10) + 5
		} else if letter == '6' {
			i = (i * 10) + 6
		} else if letter == '7' {
			i = (i * 10) + 7
		} else if letter == '8' {
			i = (i * 10) + 8
		} else if letter == '9' {
			i = (i * 10) + 9
		} else if letter == '0' {
			i = i * 10
		} else if letter == ' ' {
			return 0
		} else {
			return 0
		}
	}
	if isnegative {
		i *= -1
	}
	return i
}

func main() {
	args := os.Args
	length := 0
	for index := range args {
		length = index
	}

	for i := 1; i < length+1; i++ {
		if args[1] == "--upper" {
			val := Atoi(args[i])
			if val > 0 && val < 27 {
				newval := rune(Atoi(args[i]))
				z01.PrintRune(newval + 64)
			} else if i != 1 {
				z01.PrintRune(' ')
			}
		} else {
			val := Atoi(args[i])
			if val > 0 && val < 27 {
				newval := rune(Atoi(args[i]))
				z01.PrintRune(newval + 96)
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
