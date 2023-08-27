func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    freq := make(map[rune]int)

    for _, ru := range s {
        freq[ru]++
    }

    for _, ru := range t {
        freq[ru]--
    }

    for _, v := range freq {
        if v != 0 {
            return false
        }
    }

    return true
}

