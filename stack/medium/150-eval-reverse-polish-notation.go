import "strconv"

func evalRPN(tokens []string) int {
    operatorMap := map[string]func(a, b int)int{
        "+": func(a, b int) int { return a + b },
        "-": func(a, b int) int { return a - b },
        "*": func(a, b int) int { return a * b },
        "/": func(a, b int) int { return a / b },
    }
    
    stack := make([]int, 0)

    for _, token := range tokens {
        if operation, ok := operatorMap[token]; !ok {
            val,_ := strconv.Atoi(token)
            stack = append(stack, val)
        } else {
            // a operator b, a being first pushed item and b being second
            a, _ := strconv.Atoi(stack[len(stack) - 2])
            b, _ := strconv.Atoi(stack[len(stack) - 1])
            result := strconv.Itoa(operation(a, b))
            stack = stack[:len(stack)-2]
            stack = append(stack, result)
        }
    }
    
    res, _ := strconv.Atoi(stack[len(stack)-1])
    return res
}
