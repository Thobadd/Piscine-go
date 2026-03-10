package piscine

import "fmt"

func QuadA(x, y int) {
	// Check if dimensions are valid
	if x <= 0 || y <= 0 {
		return
	}

	// Draw each row
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Check if we're at a corner
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				print("o")
			} else if i == 0 || i == y-1 {
				// Top or bottom edge (not corners)
				print("-")
			} else if j == 0 || j == x-1 {
				// Left or right edge (not corners)
				print("|")
			} else {
				// Middle area
				print(" ")
			}
		}
		// New line after each row
		print("\n")
	}
}

// Helper function to print a single character
func print(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
	}
}
