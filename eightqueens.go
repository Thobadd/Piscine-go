package piscine

import "github.com/01-edu/z01"

func isSafe(queens [8]int, col, row int) bool {
	for i := 0; i < col; i++ {
		if queens[i] == row {
			return false
		}
		if queens[i]-row == i-col || queens[i]-row == col-i {
			return false
		}
	}
	return true
}

func printSolution(queens [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune('0' + queens[i]))
	}
	z01.PrintRune('\n')
}

func solve(queens [8]int, col int) {
	if col == 8 {
		printSolution(queens)
		return
	}
	for row := 1; row <= 8; row++ {
		if isSafe(queens, col, row) {
			queens[col] = row
			solve(queens, col+1)
		}
	}
}

func EightQueens() {
	queens := [8]int{}
	solve(queens, 0)
}
