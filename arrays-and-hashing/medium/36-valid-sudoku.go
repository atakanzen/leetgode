func isValidSudoku(board [][]byte) bool {
    hashMap := make(map[string]bool)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            curr := string(board[i][j])
            if curr == "." {
                continue
            }

            _, rowOk := hashMap[curr+"row"+string(i)]
            _, colOk := hashMap[curr+"col"+string(j)]
            _, squareOk := hashMap[curr+"square"+string(i/3) + string(j/3)]

            if rowOk || colOk || squareOk {
                return false
            }

            hashMap[curr+"row"+string(i)] = true
            hashMap[curr+"col"+string(j)] = true
            hashMap[curr+"square"+string(i/3) + string(j/3)] = true
        }
    }

    return true
}
