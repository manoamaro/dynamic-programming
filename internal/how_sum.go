package internal

func HowSum(targetSum int, numbers []int) []int {
	memo := make(map[int][]int)
	return howSum(targetSum, numbers, memo)
}

func howSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := howSum(remainder, numbers, memo)
		if remainderResult != nil {
			memo[targetSum] = append(remainderResult, num)
			return memo[targetSum]
		}
	}
	memo[targetSum] = nil
	return nil
}

func HowSumTabulation(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)

	table[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					table[i+num] = append(table[i], num)
				}
			}
		}
	}

	return table[targetSum]
}
