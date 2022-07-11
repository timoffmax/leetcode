// 9. Palindrome Number
// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
    result := false

    if x >= 0 {
        result = true
        numString := strconv.Itoa(x)
        length := len(numString)

        if length > 2 {
            isDigitsNumberEven := (0 == length % 2)

            if true == isDigitsNumberEven {
                frontIndex := 0
                middleIndex := int(length / 2) - 1
                backIndex := (length - 1)


                for i := frontIndex; i <= middleIndex ; i++ {
                    result = result && (numString[frontIndex] == numString[backIndex])

                    if false == result {
                        break;
                    }

                    frontIndex++
                    backIndex--
                }
            } else {
                frontIndex := 0
                middleIndex := int((length - 1) / 2)
                backIndex := (length - 1)

                for i := frontIndex; i < middleIndex ; i++ {
                    result = result && (numString[frontIndex] == numString[backIndex])

                    if false == result {
                        break;
                    }

                    frontIndex++
                    backIndex--
                }
            }
        } else if 2 == length {
            result = (numString[0] == numString[1])
        }
    }

    return result
}
