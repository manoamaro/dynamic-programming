package internal

func AllConstruct(target string, wordBank []string) [][]string {
	return allConstruct(target, wordBank, make(map[string][][]string))
}

func allConstruct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			suffix := target[len(word):]
			waysForRest := allConstruct(suffix, wordBank, memo)
			for _, way := range waysForRest {
				result = append(result, append([]string{word}, way...))
			}
		}
	}
	memo[target] = result
	return result
}
