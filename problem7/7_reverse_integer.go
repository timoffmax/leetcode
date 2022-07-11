// 7. Reverse Integer
// https://leetcode.com/problems/reverse-integer/

func reverseString(s string) string {
    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

func getAbsInt(i int) int {
    if i < 0 {
        i *= -1
    }

    return i
}

func getIntResultFromString(s string, isNegative bool) int {
    var result int

    result, _ = strconv.Atoi(s)

    if true == isNegative {
        result *= -1
    }

    if (result < -2147483648) || (result > 2147483647) {
        result = 0
    }

    return result
}

func reverse(x int) int {
    isNegative := x < 0
    xUnsigned := getAbsInt(x)
    xStr := strconv.Itoa(xUnsigned)
    reversedXStr := reverseString(xStr)
    result := getIntResultFromString(reversedXStr, isNegative)

    return result
}
