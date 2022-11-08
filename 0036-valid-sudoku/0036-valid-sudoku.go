func isValidSudoku(board [][]byte) bool {
	rMap := [9][9]bool{}
	cMap := [9][9]bool{}
	gMap := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val, err := strconv.Atoi(string(board[i][j]))
			if err != nil {
				continue
			}
			val--
			gIndex := i/3 + (j/3)*3
			if rMap[i][val] || cMap[j][val] || gMap[gIndex][val] {
				return false
			}
			rMap[i][val] = true
			cMap[j][val] = true
			gMap[gIndex][val] = true
		}
	}
	return true
}
