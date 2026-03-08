package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	digits := make([]int, n)
	first := true

	var generate func(start, depth int)
	generate = func(start, depth int) {
		if depth == n {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			for _, d := range digits {
				z01.PrintRune(rune('0' + d))
			}
			return
		}
		for d := start; d <= 9-(n-1-depth); d++ {
			digits[depth] = d
			generate(d+1, depth+1)
		}
	}
	generate(0, 0)
	z01.PrintRune('\n')
}
