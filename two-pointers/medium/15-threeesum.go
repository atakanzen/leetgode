import "sort"

func threeSum(nums []int) [][]int {
    res := make([][]int, 0)

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })



    for i, n := range nums {
        if i > 0 && n == nums[i - 1] {
            continue
        }

        left, right := i + 1, len(nums) - 1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum > 0 {
                right--
            } else if sum < 0 {
                left++
            } else {
                res = append(res, []int{n, nums[left], nums[right]})
                left++
                for nums[left] == nums[left-1] && left < right {
                    left++
                }
            }
        }
    }

    return res
}
