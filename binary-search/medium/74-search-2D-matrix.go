func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, (m * n)-1

    for left <= right {
        mid := left + (right - left) / 2
        midM := mid / n
        midN := mid % n
        if matrix[midM][midN] == target {
            return true
        } else if matrix[midM][midN] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return false
}
