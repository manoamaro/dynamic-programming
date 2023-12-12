package internal

func CountConstruct(target string, wordBank []string) int {
	memo := make(map[string]int)
	return countConstruct(target, wordBank, memo)
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}

	if target == "" {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			suffix := target[len(word):]
			numWaysForRest := countConstruct(suffix, wordBank, memo)
			totalCount += numWaysForRest
		}
	}

	memo[target] = totalCount
	return totalCount
}
