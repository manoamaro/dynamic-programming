package internal

func GridTravelerTabulation(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	slice2d := make([][]int, m+1)
	for i := range slice2d {
		slice2d[i] = make([]int, n+1)
	}

	slice2d[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			val := slice2d[i][j]
			if i+1 <= m {
				slice2d[i+1][j] += val
			}
			if j+1 <= n {
				slice2d[i][j+1] += val
			}
		}
	}

	return slice2d[m][n]
}
