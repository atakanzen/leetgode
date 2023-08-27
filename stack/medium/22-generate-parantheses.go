import "strings"

func generateParenthesis(n int) []string {
    stack := make([]string, 0)
    res := make([]string, 0)

    var backtrack func(openN, closeN int)
    backtrack = func(openN, closeN int) {
        if openN == closeN && openN == n {
            res = append(res, strings.Join(stack, ""))
            return
        }

        if openN < n {
            stack = append(stack, "(")
            backtrack(openN+1, closeN)
            stack = stack[:len(stack)-1]
        }

        if closeN < openN {
            stack = append(stack, ")")
            backtrack(openN, closeN+1)
            stack = stack[:len(stack)-1]
        }
    }

    backtrack(0,0)
    
    return res
}
