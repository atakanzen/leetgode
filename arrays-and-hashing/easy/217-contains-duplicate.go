func containsDuplicate(nums []int) bool {
    nums = quicksort(nums)

    for i:=0; i < len(nums) -1; i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }

    return false
}

func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left, right := 0, len(nums)-1

	pivot := rand.Intn(len(nums))
	nums[pivot], nums[right] = nums[right], nums[pivot]
	pivot = right

	for i := 0; i < pivot; i++ {
		if nums[i] < nums[pivot] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]

	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}

