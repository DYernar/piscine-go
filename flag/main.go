package main

import (
	"fmt"
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func main() {
	length := 0
	lengthOfInsert := 0
	lengthOfOrd := 0
	arguments := os.Args
	pointToStart := 0
	insert := 0
	order := 0
	var arrIns []byte
	var arrOrd []byte
	containsOrder := false
	containsInsert := false
	lengthTotal := 0
	for range arguments {
		length++
	}
	if length == 1 || arguments[1] == "--help" || arguments[1] == "-h" {
		fmt.Print("--insert\n  -i\n    This flag inserts the string into the string passed as argument.\n")
		fmt.Print("--order\n  -o\n    This flag will behave like a boolean, if it is called it will order the argument.\n")
	} else {

		for index, item := range arguments {
			if piscine.Index(item, "--i=") >= 0 {
				containsInsert = true
				insert = index
				pointToStart = 4
			} else if piscine.Index(item, "--insert=") >= 0 {
				containsInsert = true
				insert = index
				pointToStart = 9
			}
		}

		for index, item := range arguments {
			if item == "--o" || item == "--order" {
				containsOrder = true
				order = index
			}
		}

		if containsOrder {
			for _, item := range arguments[order+1] {
				arrOrd = append(arrOrd, byte(item))
			}
		}

		if containsInsert {
			for range arguments[insert] {
				lengthOfInsert++
			}
			for i := pointToStart; i < lengthOfInsert; i++ {
				arrIns = append(arrIns, arguments[insert][i])
			}

		}

		if containsOrder && containsInsert {
			for _, item := range arrOrd {
				arrIns = append(arrIns, item)
			}
			for range arrIns {
				lengthTotal++
			}
			for j := 0; j < lengthTotal; j++ {
				for i := 0; i < lengthTotal-1; i++ {
					if arrIns[i] > arrIns[i+1] {
						temp := arrIns[i]
						arrIns[i] = arrIns[i+1]
						arrIns[i+1] = temp
					}
				}
			}
			for i := 0; i < lengthTotal; i++ {
				z01.PrintRune(rune(arrIns[i]))
			}
			z01.PrintRune('\n')
		} else if containsInsert && !containsOrder {
			for i := insert + 1; i < length; i++ {
				fmt.Print(arguments[i])
			}
			for _, item := range arrIns {
				z01.PrintRune(rune(item))
			}
			z01.PrintRune('\n')
		} else if containsOrder && !containsInsert {
			for index := range arrOrd {
				lengthOfOrd = index
			}
			for j := 0; j < lengthOfOrd; j++ {
				for i := 0; i < lengthOfOrd; i++ {
					if arrOrd[i] > arrOrd[i+1] {
						temp := arrOrd[i]
						arrOrd[i] = arrOrd[i+1]
						arrOrd[i+1] = temp
					}
				}
			}
			for _, item := range arrOrd {
				z01.PrintRune(rune(item))
			}
			z01.PrintRune('\n')
		} else {
			for index, item := range arguments {
				if index > 0 {
					fmt.Print(item)
					z01.PrintRune('\n')
				}
			}
		}

	}

}
