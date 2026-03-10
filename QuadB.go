package piscine

import "fmt"

func QuadB(x, y int) {
	// Check if dimensions are valid
	if x <= 0 || y <= 0 {
		return
	}

	// Draw each row
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top row
			if i == 0 {
				if j == 0 {
					fmt.Print("/")
				} else if j == x-1 {
					fmt.Print("\\")
				} else {
					fmt.Print("*")
				}
			} else if i == y-1 { // Bottom row
				if j == 0 {
					fmt.Print("\\")
				} else if j == x-1 {
					fmt.Print("/")
				} else {
					fmt.Print("*")
				}
			} else { // Middle rows
				if j == 0 || j == x-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		// New line after each row
		fmt.Println()
	}
}
