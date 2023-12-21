package internal

func CanConstruct(target string, wordBank []string) bool {
	memo := make(map[string]bool)
	return canConstruct(target, wordBank, memo)
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			suffix := target[len(word):]
			if canConstruct(suffix, wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

func CanConstructTabulation(target string, wordBank []string) bool {
	table := make([]bool, len(target)+1)

	// base case = empty string
	table[0] = true

	for i := 0; i <= len(target); i++ {
		if table[i] {
			for _, word := range wordBank {
				// if the word matches the characters starting at position i
				if i+len(word) <= len(target) && target[i:i+len(word)] == word {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[len(target)]
}
