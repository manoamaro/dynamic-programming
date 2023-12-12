package internal

func FibTabulation(n int) int {
	if n <= 2 {
		return 1
	}

	table := make([]int, n+1)
	table[1] = 1
	table[2] = 1

	for i := 3; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}
