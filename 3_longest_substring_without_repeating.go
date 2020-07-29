// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

func lengthOfLongestSubstring(s string) int {
	result := search(s, 0)

	return result
}

func search(s string, result int) int {
	usedChars := make(map[int]string)

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		isDuplicate, duplicateIndex := inArray(char, usedChars)

		if true == isDuplicate {
			result = getGreaterNumber(result, len(usedChars))
			s = s[duplicateIndex+1:]

			return search(s, result)
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

			break
		}
	}

	return found, index
}
