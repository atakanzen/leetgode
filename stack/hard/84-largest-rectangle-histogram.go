type HistogramBar struct {
    Index int
    Height int
}

func largestRectangleArea(heights []int) int {
    stack := make([]HistogramBar, 0)
    maxArea := 0
    
    for i, h := range heights {
        start := i
        for len(stack) != 0 && stack[len(stack)-1].Height > h {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            maxArea = max(maxArea, (i - top.Index) * top.Height)
            start = top.Index
        }
        stack = append(stack, HistogramBar{Index: start, Height: h})
    }

    for _, h := range stack {
        maxArea = max(maxArea, h.Height * (len(heights) - h.Index))
    }

    return maxArea
}

func max(a, b int) int {
    if a > b {
        return a
    } 
    return b
}

