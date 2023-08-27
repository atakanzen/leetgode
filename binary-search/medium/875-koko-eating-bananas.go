func minEatingSpeed(piles []int, h int) int {
    low, high, k := 1, 0, 1
    
    for i := range piles {
        high = max(high, piles[i])
    }

    for low <= high {
        mid := low + (high - low) / 2
        
        if canEat(piles, h, mid) {
            k = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return k
}

func canEat(piles []int, timeLimit, speed int) bool {
    hoursSpent := 0
    for _, banana := range piles {
        hoursSpent += (banana + speed - 1) / speed
        if hoursSpent > timeLimit {
            return false
        }
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
