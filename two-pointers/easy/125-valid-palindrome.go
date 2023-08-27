import (
    "unicode"
)

func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    runes := []rune(s)

    for l < r {
        left := unicode.ToLower(runes[l])
        right := unicode.ToLower(runes[r])

        if !isAlphaNumeric(left) {
            l++
            continue
        }

        if !isAlphaNumeric(right) {
            r--
            continue
        }

        if left != right {
            return false
        }

        l++
        r--
    }

    return true
}   

func isAlphaNumeric(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}
