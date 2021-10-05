// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := 0.0

	allNumbers := append(nums1, nums2...)
	sort.Ints(allNumbers)

	lenght := len(allNumbers)
	middleIndex := (lenght - 1) / 2

	if 0 == (len(allNumbers) % 2) {
		medianSum := allNumbers[middleIndex] + allNumbers[middleIndex+1]
		result = float64(medianSum) / 2
	} else {
		result = float64(allNumbers[middleIndex])
	}

	return result
}
