package internal

func BestSum(targetSum int, numbers []int) []int {
	memo := make(map[int][]int)
	return bestSum(targetSum, numbers, memo)
}

func bestSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int = nil

	for _, num := range numbers {
		reminder := targetSum - num
		reminderCombination := bestSum(reminder, numbers, memo)
		if reminderCombination != nil {
			combination := append(reminderCombination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}
	memo[targetSum] = shortestCombination
	return shortestCombination
}
