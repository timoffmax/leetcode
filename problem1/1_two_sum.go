// 1. Two Sum
// https://leetcode.com/problems/two-sum/

package main

import "fmt"

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var result []int

	for i, firstNumber := range nums {
		for j, secondNumber := range nums {
			if j == i {
				continue
			}

			sum := firstNumber + secondNumber

			if sum == target {
				result = append(result, i, j)

				return result
			}
		}
	}

	return result
}
