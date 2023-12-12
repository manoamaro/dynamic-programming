package internal

func CanSum(targetSum int, numbers []int) bool {
	return canSum(targetSum, numbers, make(map[int]bool))
}

func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		remainder := targetSum - num
		if canSum(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}
