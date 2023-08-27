func longestConsecutive(nums []int) int {
    hashmap := make(map[int]bool)
    longest := 0

    for _, n := range nums {
        // Add distinct values only
        if _, ok := hashmap[n]; !ok {
            hashmap[n] = true
        }
    }

    for n, _ := range hashmap {
        // Start of sequence
        if _, ok := hashmap[n-1]; !ok {
            count := 0
            for _, ok := hashmap[n+count]; ok; _, ok =hashmap[n+count] {
                count++
            }
            if count > longest {
                longest = count
            }
        }    
    }

    return longest
}
