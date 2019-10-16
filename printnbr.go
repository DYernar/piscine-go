package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var isnegative bool
	var strVal string
	if n > 0 {
		isnegative = false
		for i := 0; i <= 10; i++ {
			if n != 0 {
			strVal = string((n%10)+48) + strVal
			n =(n - (n%10))/10
			}
		}
	} else if n < 0 {
		isnegative = true
		n = -n
		for i := 0; i <= 10; i++ {
			if n != 0 {
			strVal = string((n%10)+48) + strVal
			n =(n - (n%10))/10
			}
		}
	} else {
		strVal = "0"
	}
	if !isnegative {
		for j := 0; j <len(strVal); j++ {
			z01.PrintRune(rune(strVal[j]))
		}
	} else {
		strVal = "-" + strVal
		for j := 0; j <len(strVal); j++ {
			z01.PrintRune(rune(strVal[j]))
		}
	}
}