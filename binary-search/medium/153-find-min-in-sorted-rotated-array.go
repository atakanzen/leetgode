func findMin(nums []int) int {
    minIndex := -1
    low, hi := 0, len(nums)-1
    lastItem := nums[hi]

    for low <= hi {
        mid := low + (hi - low) / 2
        if nums[mid] <= lastItem {
            minIndex = mid
            hi = mid - 1
        } else {
            low = mid + 1
        }
    }

    return nums[minIndex]
}
