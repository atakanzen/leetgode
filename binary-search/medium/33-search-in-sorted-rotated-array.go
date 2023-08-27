func search(nums []int, target int) int {
    index := -1
    low, hi := 0, len(nums)-1

    for low <= hi {
        mid := low + (hi - low) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[low] <= nums[mid] {
            if target > nums[mid] || target < nums[low] {
                low = mid + 1
            } else {
                hi = mid - 1
            }
        } else {
            if target < nums[mid] || target > nums[hi] {
                hi = mid - 1
            } else {
                low = mid + 1
            }
        }
    }

    return index
}
