package piscine

import "fmt"

func QuadC(x, y int) {
	// Check if dimensions are valid
	if x <= 0 || y <= 0 {
		return
	}

	// Draw each row
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top row
			if i == 0 {
				if j == 0 || j == x-1 {
					fmt.Print("A")
				} else {
					fmt.Print("B")
				}
			} else if i == y-1 { // Bottom row
				if j == 0 || j == x-1 {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			} else { // Middle rows
				if j == 0 || j == x-1 {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
		}
		// New line after each row
		fmt.Println()
	}
}
