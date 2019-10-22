package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n > 0 {
		var length int
		var arr []rune
		for i := 0; n != 0; i++ {
			arr = append(arr, rune(n%10+48))
			n = n / 10
			length++
		}
		for index := range arr {
			for j := range arr {
				if arr[index] > arr[j] {
					arr[index], arr[j] = arr[j], arr[index]
				}
			}
		}
		for j := length - 1; j >= 0; j-- {
			z01.PrintRune(arr[j])
		}
	}

}
