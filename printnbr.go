package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var strVal string
	var length int
	var isnegative bool
	if n > 0 {
		isnegative = false
		for i := 0; i <= 10; i++ {
			if n != 0 {
				if n == 1 {
					strVal = "1" + strVal
				} else if n % 10 == 2 {
					strVal = "2" + strVal
				} else if n % 10 == 3 {
					strVal = "3" + strVal
				} else if n % 10 == 4  {
					strVal = "4" + strVal
				} else if n % 10 == 5 {
					strVal = "5" + strVal
				} else if n % 10 == 6 {
					strVal = "6" + strVal
				} else if n % 10 == 7 {
					strVal = "7" + strVal
				} else if n % 10 == 8 {
					strVal = "8" + strVal
				} else {
					strVal = "9" + strVal
				}
				n = (n - (n % 10)) / 10
				length++
			}
		}
	} else if n < 0 {
		isnegative = true
		n = -n
		for i := 0; i <= 10; i++ {
			if n != 0 {
				if n == 1 {
					strVal = "1" + strVal
				} else if n % 10 == 2 {
					strVal = "2" + strVal
				} else if n % 10 == 3 {
					strVal = "3" + strVal
				} else if n % 10 == 4  {
					strVal = "4" + strVal
				} else if n % 10 == 5 {
					strVal = "5" + strVal
				} else if n % 10 == 6 {
					strVal = "6" + strVal
				} else if n % 10 == 7 {
					strVal = "7" + strVal
				} else if n % 10 == 8 {
					strVal = "8" + strVal
				} else {
					strVal = "9" + strVal
				}
				n = (n - (n % 10)) / 10
				length++
			}
		}
	}
	if isnegative {
		strVal = "-" + strVal
		length++
	} else if !isnegative {

	} else {
		strVal = "0"
		length = 1
	}
	for j := 0; j <length; j++ {
		z01.PrintRune(rune(strVal[j]));
	}
}
