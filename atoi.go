package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	negative := false
	start := 0

	if s[0] == '+' || s[0] == '-' {
		if len(s) == 1 {
			return 0
		}
		negative = s[0] == '-'
		start = 1
	}

	for _, ch := range s[start:] {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}

	if negative {
		return -result
	}
	return result
}
