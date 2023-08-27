func trap(height []int) int {
    water := 0
    l, r := 0, len(height)-1
    maxLeft, maxRight := height[l], height[r]

    for l < r {
        if maxLeft < maxRight {
            l++
            maxLeft = max(maxLeft, height[l])
            water += maxLeft - height[l]
        } else {
            r--
            maxRight = max(maxRight, height[r])
            water += maxRight - height[r]
        }
    }

    return water
}   

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
