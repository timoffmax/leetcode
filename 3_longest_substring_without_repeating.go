// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

func lengthOfLongestSubstring(s string) int {
	result := search(s, 0)

	return result
}

func search(s string, startFrom int) int {
	result := 0
	usedChars := make(map[int]string)

	for i := startFrom; i < len(s); i++ {
		char := string(s[i])
		isDuplicate, duplicateIndex := inArray(char, usedChars)

		if true == isDuplicate {
			result = getGreaterNumber(result, len(usedChars))
			revisedResult := search(s, duplicateIndex+1)

			if revisedResult > result {
				result = revisedResult
			}

			usedChars = make(map[int]string)
		}

		usedChars[i] = char
	}

	result = getGreaterNumber(result, len(usedChars))

	return result
}

func getGreaterNumber(prevValue int, currentValue int) int {
	result := prevValue

	if currentValue > prevValue {
		result = currentValue
	}

	return result
}

func inArray(needle string, haystack map[int]string) (bool, int) {
	found := false
	index := 0

	for i, item := range haystack {
		if item == needle {
			found = true
			index = i
		}
	}

	return found, index
}
