func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int
        for _, c := range s {
            count[c - 'a']++
        }
        freqMap[count] = append(freqMap[count], s)
    }

    result := make([][]string, len(freqMap))
    i := 0
    for _, group := range freqMap {
        result[i] = group
        i++
    }

    return result
}
