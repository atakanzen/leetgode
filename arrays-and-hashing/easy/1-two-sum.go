func twoSum(nums []int, target int) []int {
    prev := make(map[int]int) // val : index

    for i, n := range nums {
        diff := target - n
        if index, ok := prev[diff]; ok {
            return []int{index, i}
        }
        prev[n] = i
    }

    return nil
}
