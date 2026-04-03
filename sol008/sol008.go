package sol008
func IsValidSudoku(board [][]byte) bool {
   rows := make([]int, 9)
   cols := make([]int, 9)
   sqrs := make([]int, 9)
   
   for r := 0; r < 9; r++{
	for c := 0; c < 9; c++{
		if board[r][c] == '.'{
			continue
		}

		val := board[r][c] - '1'
		bit  := 1 << val
		sqrcIdx := (r/3)*3 + (c/3)

		if rows[r]&bit != 0 || cols[c]&bit != 0 || sqrs[sqrcIdx]&bit != 0{
			return  false
		}

		rows[r] |= bit
		cols[c] |= bit
		sqrs[sqrcIdx] |= bit
	}
   }

   return true
}