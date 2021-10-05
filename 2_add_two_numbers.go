// 2. Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum big.Int
	number1 := getNumber(l1)
	number2 := getNumber(l2)
	sum.Add(&number1, &number2)

	answer := prepareAnswer(sum)

	return &answer
}

func getNumber(node *ListNode) big.Int {
	var result big.Int

	numbersArray := nodeToArray(node)
	numbersString := arrayToReverseNumber(numbersArray)
	result.SetString(numbersString, 10)

	return result
}

func nodeToArray(node *ListNode) []int {
	var result []int
	currentNode := node

	for {
		result = append(result, currentNode.Val)

		if nil == currentNode.Next {
			break
		} else {
			currentNode = currentNode.Next
		}
	}

	return result
}

func prepareAnswer(number big.Int) ListNode {
	numbersString := number.String()
	numbers := stringToArray(numbersString)
	reverseArray := reverseArray(numbers)
	result := arrayToListNode(reverseArray)

	return result
}

func arrayToReverseNumber(numbers []int) string {
	reversedArray := reverseArray(numbers)
	result := arrayToString(reversedArray)

	return result
}

func arrayToString(a []int) string {
	result := strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")

	return result
}

func stringToArray(s string) []int {
	var result []int
	arr := strings.Split(s, "")

	for i := range arr {
		number, _ := strconv.ParseInt(arr[i], 10, 64)
		result = append(result, int(number))
	}

	return result
}

func reverseArray(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

func arrayToListNode(numbers []int) ListNode {
	result := ListNode{}

	var nodes []ListNode
	var currentNode *ListNode

	for _, number := range numbers {
		node := ListNode{
			Val: number,
		}

		nodes = append(nodes, node)
	}

	for i, node := range nodes {
		if (i + 1) < len(nodes) {
			node.Next = &nodes[i+1]
		}

		if nil == currentNode {
			currentNode = &result
		} else {
			currentNode = currentNode.Next

		}

		currentNode.Val = node.Val
		currentNode.Next = node.Next
	}

	return result
}
