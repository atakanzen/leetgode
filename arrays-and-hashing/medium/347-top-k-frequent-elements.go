func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
    freq := make([][]int, len(nums)+1)

    for _, n := range nums {
        if _, ok := count[n]; !ok {
            count[n] = 1
        } else {
            count[n]++
        }
    }

    for n, c := range count {
        freq[c] = append(freq[c], n)
    }

    res := make([]int, 0)
    for i := len(freq)-1; i > 0; i-- {
        if len(freq[i]) != 0 {
            res = append(res, freq[i]...)
        }
    }

    return res[:k]
}
