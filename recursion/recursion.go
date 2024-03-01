package recursion

func Fatorial(n int) int {
	switch {
	case n < 0:
		return 0
	case n > 0:
		return Fatorial(n-1) * n
	default:
		return 1
	}
}
