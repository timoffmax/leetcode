// 8. String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/

func clampInt(i int64) int {
    var result int

    if i < -2147483648 {
        result = -2147483648
    } else if i > 2147483647 {
        result = 2147483647
    } else {
        result = int(i)
    }

    return result
}

func myAtoi(s string) int {
    result := 0
    trimmedValue := strings.TrimSpace(s)

    if len(trimmedValue) > 0 {
        firstChar := trimmedValue[0:1]
        isNegative := ("-" == firstChar)

        var unsignedValue string
        digits := ""

        if (true == isNegative) || ("+" == firstChar) {
            unsignedValue = trimmedValue[1:]
        } else {
            unsignedValue = trimmedValue
        }

        for _, c := range unsignedValue {
            if false == unicode.IsDigit(rune(c)) {
                break
            }

            digits += string(c)
        }

        if len(digits) > 0 {
            intValue, _ := strconv.ParseInt(digits, 10, 64)

            if true == isNegative {
                intValue *= -1
            }

            result = clampInt(intValue)
        } else {
            result = 0;
        }
    }

    return result
}
