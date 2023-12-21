package internal

func MinimizeSteps(k int) int {
	if k == 1 {
		return 1
	}
	c := 0

	if k%2 == 0 {
		c = 1 + MinimizeSteps(k/2)
	} else {
		c = 1 + MinimizeSteps(k-1)
	}

	return c
}

func MinimizeStepsTabulation(k int) int {
	dp := make([]int, k+1)

	dp[1] = 1
	for i := 1; i <= k; i++ {
		if i%2 == 0 {
			dp[i] = 1 + dp[i/2]
		} else {
			dp[i] = 1 + dp[i-1]
		}
	}

	return dp[k]
}
