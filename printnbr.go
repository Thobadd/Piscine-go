package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		printNeg(n)
	} else {
		printPos(n)
	}
}

func printPos(n int) {
	if n >= 10 {
		printPos(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}

func printNeg(neg int) {
	if neg <= -10 {
		printNeg(neg / 10)
	}
	z01.PrintRune(rune('0' - neg%10))
}
