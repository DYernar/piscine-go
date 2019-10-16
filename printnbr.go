package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var strVal string
	var length int
	var isnegative string
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
	} else {
		if n > 0 {
			isnegative = "f"
			for i := 0; i <= 18; i++ {
				if n != 0 {
					if n%10 == 1 {
						strVal = "1" + strVal
					} else if n%10 == 2 {
						strVal = "2" + strVal
					} else if n%10 == 3 {
						strVal = "3" + strVal
					} else if n%10 == 4 {
						strVal = "4" + strVal
					} else if n%10 == 5 {
						strVal = "5" + strVal
					} else if n%10 == 6 {
						strVal = "6" + strVal
					} else if n%10 == 7 {
						strVal = "7" + strVal
					} else if n%10 == 8 {
						strVal = "8" + strVal
					} else if n%10 == 0 {
						strVal = "0" + strVal
					} else {
						strVal = "9" + strVal
					}
					n = (n - (n % 10)) / 10
					length++
				}
			}
		} else if n < 0 {
			isnegative = "t"
			n = -n
			for i := 0; i <= 18; i++ {
				if n != 0 {
					if n%10 == 1 {
						strVal = "1" + strVal
					} else if n%10 == 2 {
						strVal = "2" + strVal
					} else if n%10 == 3 {
						strVal = "3" + strVal
					} else if n%10 == 4 {
						strVal = "4" + strVal
					} else if n%10 == 5 {
						strVal = "5" + strVal
					} else if n%10 == 6 {
						strVal = "6" + strVal
					} else if n%10 == 7 {
						strVal = "7" + strVal
					} else if n%10 == 8 {
						strVal = "8" + strVal
					} else if n%10 == 0 {
						strVal = "0" + strVal
					} else {
						strVal = "9" + strVal
					}
					n = (n - (n % 10)) / 10
					length++
				}
			}
		} else if n == 0 {
			isnegative = "z"
		}
		if isnegative == "t" {
			strVal = "-" + strVal
			length++
		} else if isnegative == "f" {

		} else {
			strVal = "0"
			length = 1
		}
		for j := 0; j < length; j++ {
			if strVal[j] == '1' {
				z01.PrintRune('1')
			} else if strVal[j] == '2' {
				z01.PrintRune('2')
			} else if strVal[j] == '3' {
				z01.PrintRune('3')
			} else if strVal[j] == '4' {
				z01.PrintRune('4')
			} else if strVal[j] == '5' {
				z01.PrintRune('5')
			} else if strVal[j] == '6' {
				z01.PrintRune('6')
			} else if strVal[j] == '7' {
				z01.PrintRune('7')
			} else if strVal[j] == '8' {
				z01.PrintRune('8')
			} else if strVal[j] == '9' {
				z01.PrintRune('9')
			} else if strVal[j] == '0' {
				z01.PrintRune('0')
			} else {
				z01.PrintRune('-')
			}
		}
	}
}
