// 5. Longest Palindromic String
// https://leetcode.com/problems/longest-palindromic-substring/

package main

import "fmt"

func main() {
	result := longestPalindrome("babad")
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		for j := len(s); j >= i; j-- {
			var subStr string

			if j == i {
				subStr = string(s[j])
			} else {
				subStr = s[i:j]
			}

			if true == isPalindrome(subStr) {
				if len(subStr) > len(result) {
					result = subStr
				}
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	var result bool

	strLen := len(s)
	last := strLen - 1

	if 1 == strLen {
		return true
	}

	for start, end := 0, last; start <= end; start, end = start+1, end-1 {
		result = (s[start] == s[end])

		if false == result {
			break
		}
	}

	return result
}
