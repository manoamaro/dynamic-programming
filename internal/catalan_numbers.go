package internal

func CatalanNumbers(n int) int {
	if n <= 1 {
		return 1
	}
	a := (4 * n * n) - 2*n
	b := n*n + n
	return (a * CatalanNumbers(n-1)) / b
}
