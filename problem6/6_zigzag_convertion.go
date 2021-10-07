// 6. ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"fmt"
	"strings"
)

func main() {
	result := convert("PAYPALISHIRING", 3)
	fmt.Println(result)

	result = convert("PAYPALISHIRING", 4)
	fmt.Println(result)

	result = convert("AB", 1)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	var result string

	maxFullColumns := (len(s) / numRows) - numRows
	singleCharColumns := len(s) - maxFullColumns
	numColumns := maxFullColumns + singleCharColumns
	rowIdx := 0
	columnIdx := 0
	isBackwardDirection := false
	resultArr := createStringSlice(numRows, numColumns)

	for i := 0; i < len(s); i++ {
		for row := 0; row < numRows; row++ {
			if row == rowIdx {
				resultArr[row][columnIdx] = string(s[i])

				if true == isBackwardDirection {
					rowIdx--
				} else {
					rowIdx++
				}

				break
			}
		}

		if numRows > 1 {
			if numRows == rowIdx {
				columnIdx++
				isBackwardDirection = true
				rowIdx -= 2
			} else if -1 == rowIdx {
				columnIdx++
				isBackwardDirection = false
				rowIdx += 2
			}
		} else {
			columnIdx++
			rowIdx--
		}
	}

	result = stringifyResult(resultArr)

	return result
}

func createStringSlice(numRows int, numColumns int) [][]string {
	s := make([][]string, numRows)

	for i := range s {
		s[i] = make([]string, numColumns)
	}

	return s
}

func stringifyResult(a [][]string) string {
	result := ""

	for _, row := range a {
		stringRow := arrayToString(row)

		if "" != stringRow {
			result += stringRow
		} else {
			break
		}
	}

	return result
}

func arrayToString(a []string) string {
	result := strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")

	return result
}
