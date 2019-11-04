package main

import (
	"bufio"
	"fmt"
	"os"
)

func Raid1a(x, y int) []rune {
	var arr []rune
	if x > 0 {
		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				for j := 0; j < x; j++ {
					if (j == 0) || (j == x-1) {
						arr = append(arr, 'o')
					} else {
						arr = append(arr, '-')
					}
				}
				arr = append(arr, '\n')
			} else if i != 0 || i != y-1 {
				for j := 0; j < x; j++ {
					if (j == 0) || (j == x-1) {
						arr = append(arr, '|')
					} else {
						arr = append(arr, ' ')
					}
				}
				arr = append(arr, '\n')
			}
		}
	}
	return arr
}

func Raid1b(x, y int) []rune {
	var arr []rune
	if x > 0 {
		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				for j := 0; j < x; j++ {
					if i == 0 && j == 0 {
						arr = append(arr, '/')
					} else if i == 0 && j == x-1 {
						arr = append(arr, '\\')
					} else if i == y-1 && j == 0 {
						arr = append(arr, '\\')
					} else if i == y-1 && j == x-1 {
						arr = append(arr, '/')
					} else {
						arr = append(arr, '*')
					}
				}
			} else {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, '*')
					} else {
						arr = append(arr, ' ')
					}
				}
			}
			arr = append(arr, '\n')
		}
	}
	return arr
}

func Raid1c(x, y int) []rune {
	var arr []rune
	if x > 0 {
		for i := 0; i < y; i++ {
			if i == 0 {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, 'A')
					} else {
						arr = append(arr, 'B')
					}
				}
			} else if i == y-1 {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, 'C')
					} else {
						arr = append(arr, 'B')
					}
				}
			} else {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, 'B')
					} else {
						arr = append(arr, ' ')
					}
				}
			}
			arr = append(arr, '\n')
		}
	}
	return arr
}
func Raid1d(x, y int) []rune {
	var arr []rune
	if x > 0 {
		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				for j := 0; j < x; j++ {
					if i == 0 && j == 0 {
						arr = append(arr, 'A')
					} else if i == 0 && j == x-1 {
						arr = append(arr, 'C')
					} else if i == y-1 && j == 0 {
						arr = append(arr, 'A')
					} else if i == y-1 && j == x-1 {
						arr = append(arr, 'C')
					} else {
						arr = append(arr, 'B')
					}
				}
			} else {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, 'B')
					} else {
						arr = append(arr, ' ')
					}
				}
			}
			arr = append(arr, '\n')
		}
	}
	return arr
}

func Raid1e(x, y int) []rune {
	var arr []rune
	if x > 0 {
		for i := 0; i < y; i++ {
			if i == 0 || i == y-1 {
				for j := 0; j < x; j++ {
					if i == 0 && j == 0 {
						arr = append(arr, 'A')
					} else if i == 0 && j == x-1 {
						arr = append(arr, 'C')
					} else if i == y-1 && j == 0 {
						arr = append(arr, 'C')
					} else if i == y-1 && j == x-1 {
						arr = append(arr, 'A')
					} else {
						arr = append(arr, 'B')
					}
				}
			} else {
				for j := 0; j < x; j++ {
					if j == 0 || j == x-1 {
						arr = append(arr, 'B')
					} else {
						arr = append(arr, ' ')
					}
				}
			}
			arr = append(arr, '\n')
		}
	}
	return arr
}Raid1a
Raid1a
func inputCheck(arr Raid1a[]rune) bool {
	for _, item := rRaid1aange arr {
		switch item {
		case 'A':
		case 'B':
		case 'C':
		case 'o':
		case '\\':
		case '-':
		case '|':
		case '/':
		case '*':
		case ' ':
		default:
			return false
		}
	}
	return true
}

func notRaid1() {
	fmt.Println("Not a Raid function")
	os.Exit(0)
}
func compareArrays(redFile []rune, raid1Arr []rune) bool {
	for index, item := range raid1Arr {
		if redFile[index] != item {
			return false
		}
	}
	return true
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		text, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, text)
	}

	if inputCheck(output) {
		notRaid1()
	}
	total := 0
	rows := 0
	colomns := 0
	for _, item := range output {
		if item == '\n' {
			rows++
			total--
		}
		total++
	}
	colomns = total / rows
	raid1a := Raid1a(colomns, rows)
	raid1b := Raid1b(colomns, rows)
	raid1c := Raid1c(colomns, rows)
	raid1d := Raid1d(colomns, rows)
	raid1e := Raid1e(colomns, rows)
	isRaid1a := compareArrays(output, raid1a)
	isRaid1b := compareArrays(output, raid1b)
	isRaid1c := compareArrays(output, raid1c)
	isRaid1d := compareArrays(output, raid1d)
	isRaid1e := compareArrays(output, raid1e)
	var finalStr [10]string
	j := 0
	if isRaid1a {
		finalStr[j] = "[raid1a]"
		j++
	} else {
		finalStr[j] = ""
		j++
	}
	if isRaid1b {
		finalStr[j] = "[raid1b]"
		j++
	} else {
		finalStr[j] = ""
		j++
	}
	if isRaid1c {
		finalStr[j] = "[raid1c]"
		j++
	} else {
		finalStr[j] = ""
		j++
	}
	if isRaid1d {
		finalStr[j] = "[raid1d]"
		j++
	} else {
		finalStr[j] = ""
		j++
	}
	if isRaid1e {
		finalStr[j] = "[raid1e]"
	} else {
		finalStr[j] = ""
	}
	printed := 0
	for _, item := range finalStr {
		if item != "" {
			if printed == 0 {
				fmt.Print(item + " [")
				fmt.Print(colomns)
				fmt.Print("] [")
				fmt.Print(rows)
				fmt.Print("]")
				printed++
			} else {
				fmt.Print(" || ")
				fmt.Print(item + " [")
				fmt.Print(colomns)
				fmt.Print("] [")
				fmt.Print(rows)
				fmt.Print("]")
			}
		}
	}
	fmt.Println()
}
