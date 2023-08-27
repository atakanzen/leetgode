import "sort"

func carFleet(target int, position []int, speed []int) int {
	// (Position, Speed)
	n := len(position)
	cars := make([][]int, n)
	for i := range position {
		cars[i] = []int{position[i], speed[i]}
	}

	sort.Slice(cars, func(a, b int) bool {
		// Sorted in ascending order based on position
		return cars[a][0] < cars[b][0]
	})

	// Time to target stack
	stack := make([]float32, 0)
	for i := n - 1; i >= 0; i-- {
		stack = append(stack, float32(target-cars[i][0])/float32(cars[i][1]))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}

