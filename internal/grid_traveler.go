package internal

import "fmt"

func GridTraveler(m, n int) int {
	memo := make(map[string]int)
	return gridTraveler(m, n, memo)
}

func gridTraveler(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d,%d,%d", m, n, n, m)
	if val, ok := memo[key]; ok {
		return val
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = gridTraveler(m-1, n, memo) + gridTraveler(m, n-1, memo)
	return memo[key]
}

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
