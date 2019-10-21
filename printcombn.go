package piscine

import (
	"github.com/01-edu/z01"
)

func verification(arr []rune) bool {
	if arr[len(arr)-1] == '9' {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i]+1 != arr[i+1] {
				return false
			}
		}
		return true
	}
	return false
}

func recursion(arr []rune) []rune {
	if arr[len(arr)-1] == '9' {
		x := '0'
		arr = arr[:len(arr)-1]
		val := recursion(arr)
		val = append(val, x)
		return val
	}
	arr[len(arr)-1]++
	return arr

}

func isLast(arr []rune) bool {
	if verification(arr) {
		return true
	}
	return false
}

func PrintCombN(n int) {
	a := '0'
	m := make([]rune, n)
	for i := 0; i < n; i++ {
		m[i] = a
		a++
	}
	for !verification(m) {
		for j := 0; j < len(m); j++ {
			z01.PrintRune(m[j])
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
		m = recursion(m)
	}
	for j := 0; j < len(m); j++ {
		z01.PrintRune(m[j])
	}

	z01.PrintRune(10)
}
