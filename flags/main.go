package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func StrLen(str string) int {
	var i int
	for range str {
		i++
	}
	return i
}

func FirstRune(s string) rune {
	var arr rune
	for index, item := range s {
		if index == 0 {
			arr = item
		}
	}
	return arr
}

func Index(s string, toFind string) int {
	if StrLen(s) == 0 {
		return -1
	}
	if StrLen(toFind) == 0 {
		return 0
	}
	strrun := []rune(s)
	run := []rune(toFind)
	var res bool
	run1 := FirstRune(toFind)
	var result int
	for index, item := range s {
		if item == run1 {
			result = index
			res = true
			for i := range run {
				if (result + i) < StrLen(s) {
					if (StrLen(s) - result) >= StrLen(toFind) {
						if run[i] != strrun[result+i] {
							res = false
						}
					} else {
						return -1
					}

				} else {
					return -1
				}

			}
			if res == true {
				return result
			}
		}
	}
	if res == true {
		return result
	} else {
		return -1
	}

}

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
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	} else {

		for index, item := range arguments {
			if Index(item, "-i=") >= 0 {
				containsInsert = true
				insert = index
				pointToStart = 3
			} else if Index(item, "--insert=") >= 0 {
				containsInsert = true
				insert = index
				pointToStart = 9
			}
		}

		for index, item := range arguments {
			if item == "-o" || item == "--order" {
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
