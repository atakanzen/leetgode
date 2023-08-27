type TemperatureItem struct {
    Index int
    Temperature int
}

func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := make([]TemperatureItem, 0)

    for i, t := range temperatures {
        for len(stack) != 0 && t > stack[len(stack)-1].Temperature {
            lastColderTemperature := stack[len(stack)-1]
            res[lastColderTemperature.Index] = i - lastColderTemperature.Index
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, TemperatureItem{Index: i, Temperature: t})
    }

    return res
}
